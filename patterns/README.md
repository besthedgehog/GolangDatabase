# Паттерны проектирования

## Woker Pool

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Задача: URL, который нужно проверить
type Job struct {
	URL string
}

// Результат: Что мы узнали об этом сайте
type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration // Как долго грузился
	Err        error         // Если сайт вообще недоступен
}

// Воркер: "Инспектор сайтов"
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Создаем HTTP клиент с таймаутом (важно для реальных задач!)
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for job := range jobs {
		// fmt.Printf("Воркер %d проверяет %s\n", id, job.URL) // (Опционально: логирование)

		start := time.Now()
		resp, err := client.Get(job.URL)
		duration := time.Since(start)

		if err != nil {
			// Если ошибка (например, нет интернета или домен не существует)
			results <- Result{
				URL: job.URL,
				Err: err,
			}
			continue
		}

		// Если успех
		results <- Result{
			URL:        job.URL,
			StatusCode: resp.StatusCode,
			Duration:   duration,
			Err:        nil,
		}
		resp.Body.Close() // Не забываем закрывать тело ответа!
	}
}

func main() {
	// Список сайтов для проверки (можно представить, что их тут 1000)
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://non-existent-website-blabla.com", // Специально сломанный
		"https://microsoft.com",
	}

	const numWorkers = 3 // 3 инспектора

	// Каналы
	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	// 1. Запуск воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. Раздача задач
	for _, url := range urls {
		jobs <- Job{URL: url}
	}
	close(jobs) // Задач больше не будет

	// 3. Фоновый "ждун" (чтобы не было дедлока!)
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Сбор и анализ результатов (Main работает тут)
	var successCount, failCount int

	fmt.Println("--- Начало проверки ---")

	// Читаем, пока канал results не закроется
	for res := range results {
		if res.Err != nil {
			fmt.Printf("[FAIL] %s | Ошибка: %v\n", res.URL, res.Err)
			failCount++
		} else {
			fmt.Printf("[OK]   %s | Код: %d | Время: %v\n", res.URL, res.StatusCode, res.Duration)
			successCount++
		}
	}

	fmt.Println("--- Итоги ---")
	fmt.Printf("Успешно: %d, Ошибок: %d\n", successCount, failCount)
}
```

## Fan in Fan out

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура данных, с которой будем работать
type User struct {
	ID       int
	Name     string
	Email    string
	Location string
}

// Эмуляция базы данных: просто возвращает список пользователей
func fetchUsers() []User {
	return []User{
		{1, "Alice", "alice@example.com", "New York"},
		{2, "Bob", "bob@example.com", "San Francisco"},
		{3, "Charlie", "charlie@example.com", "Chicago"},
		{4, "Dave", "dave@example.com", "Austin"},
		{5, "Eve", "eve@example.com", "Seattle"},
	}
}

// Функция-рабочий. Имитирует тяжелую работу (задержка 100мс)
func enrichUserData(user User) User {
	time.Sleep(100 * time.Millisecond)
	user.Location = user.Location + ", USA" // Добавляем страну
	return user
}

// ---------------------------------------------------------
// STEP 1: FAN-OUT (Раздача работы)
// Принимает: список всех пользователей и количество желаемых работников.
// Возвращает: список каналов (по одному каналу на работника).
// ---------------------------------------------------------
func fanOut(users []User, workerCount int) []<-chan User {
	// Создаем "коробку" для каналов
	channels := make([]<-chan User, workerCount)

	for i := 0; i < workerCount; i++ {
		// Создаем персональный канал для работника
		ch := make(chan User)
		channels[i] = ch

		// Запускаем работника в фоне
		go func(myChan chan User, workerID int) {
			// ВАЖНО: Закрываем канал, когда работа закончена,
			// чтобы fanIn на той стороне понял, что данных больше не будет.
			defer close(myChan)

			// Хитрый алгоритм распределения:
			// Работник 0 берет индексы: 0, 3, 6...
			// Работник 1 берет индексы: 1, 4, 7...
			// Это позволяет не использовать общий канал-очередь здесь.
			for j := workerID; j < len(users); j += workerCount {
				enriched := enrichUserData(users[j])

				// Отправляем результат в СВОЙ канал
				// Операция блокирующая, пока fanIn не прочитает.
				myChan <- enriched
			}
		}(ch, i)
	}

	return channels
}

// ---------------------------------------------------------
// STEP 2: FAN-IN (Сбор результатов)
// Принимает: кучу каналов от разных работников.
// Возвращает: ОДИН общий канал со всеми результатами.
// ---------------------------------------------------------
func fanIn(channels []<-chan User) <-chan User {
	// Создаем единый канал для слива результатов
	resultChan := make(chan User)
	var wg sync.WaitGroup

	// Функция-пересыльщик (Multiplexer).
	// Она читает из конкретного канала и перекладывает в общий.
	multiplex := func(ch <-chan User) {
		defer wg.Done()
		// Читаем, пока канал ch не закроется (в функции fanOut)
		for user := range ch {
			resultChan <- user // Перекладываем в общую трубу
		}
	}

	// Для каждого входящего канала запускаем своего пересыльщика
	for _, ch := range channels {
		wg.Add(1)
		go multiplex(ch)
	}

	// ---------------------------------------------------------
	// ВАЖНЫЙ МОМЕНТ: "Наблюдатель"
	// Мы запускаем отдельную горутину, которая ждет окончания всех пересыльщиков.
	// Если бы мы сделали wg.Wait() здесь без go func, мы бы поймали Deadlock,
	// так как Main еще не начал читать из resultChan.
	// ---------------------------------------------------------
	go func() {
		wg.Wait()          // Ждем, пока все workers закончат и закроют свои каналы
		close(resultChan)  // Закрываем общий канал. Это сигнал для Main, что всё готово.
	}()

	return resultChan
}

func main() {
	users := fetchUsers()
	workerCount := 3 // Хотим 3 параллельных потока

	start := time.Now()

	// 1. Запускаем раздачу (функция возвращает управление мгновенно)
	// Сами воркеры уже начали шуршать в фоне.
	workerChannels := fanOut(users, workerCount)

	// 2. Запускаем сборку (функция возвращает управление мгновенно)
	// Пересыльщики начали ждать данные.
	resultStream := fanIn(workerChannels)

	// 3. Читаем результаты.
	// Main блокируется здесь и вычитывает данные по мере их поступления.
	// Цикл закончится только когда "Наблюдатель" внутри fanIn закроет канал.
	processedUsers := []User{}
	for user := range resultStream {
		processedUsers = append(processedUsers, user)
		fmt.Printf("Получен пользователь: %s\n", user.Name)
	}

	elapsed := time.Since(start)
	fmt.Printf("Обработано %d пользователей за %v\n", len(processedUsers), elapsed)
}
```
