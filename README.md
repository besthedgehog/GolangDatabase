# Моя база знаний по Golang

# Оглавление
- [Базовые вещи](#базовые-вещи)
  - [Присваивание](#присваивание)
  - [Как узнать тип переменной?](#как-узнать-тип-переменной)
  - [Циклы](#циклы)
  - [Switch](#switch)
  - [Go mod init](#go-mod-init)
  - [return](#return)
- [Считывание значений из stdin](#считывание-значений-из-stdin)
- [Импорты](#импорты)
- [Директории и переменные окружения](#директории-и-переменные-окружения)
  - [Платформы и кросс-компилляция](#платформы-и-кросс-компилляция)
- [Распаковка](#распаковка)
- [defer](#defer)
  - [defer не выполняетя при os.Exit()](#defer-не-выполняетя-при-osexit)
- [ООП в golang](#ооп-в-golang)
  - [1. Наследование](#1-наследование)
  - [2. Инкапсуляция](#2-инкапсуляция)
  - [3. Полиморфизм](#3-полиморфизм)
  - [Встраивания (Эмбединги)](#встраивания-эмбединги)
    - [Как понять когда писать с именем, когда без:](#как-понять-когда-писать-с-именем-когда-без)
  - [Ещё примеры встраивания](#ещ-примеры-встраивания)
- [Интерфейсы в Golang](#интерфейсы-в-golang)
  - [Еще одно объяснение интефейсов](#еще-одно-объяснение-интефейсов)
  - [Использование интерфейсов](#использование-интерфейсов)
    - [Как можно использовать интерфейсы в тестах?](#как-можно-использовать-интерфейсы-в-тестах)
- [Паттерны проектирования](#паттерны-проектирования)
  - [Фабрика](#фабрика)
  - [Адаптер](#адаптер)
    - [Когда применять:](#когда-применять)
  - [Dependency Injection (Внедрение зависимостей)](#dependency-injection-внедрение-зависимостей)
    - [Алгоритм создания di](#алгоритм-создания-di)
- [Перенаправления методов (перегрузка)](#перенаправления-методов-перегрузка)
- [Замыкания](#замыкания)
- [Generic](#generic)
- [Resful API](#resful-api)
    - [Почему это круто](#почему-это-круто)
- [Указатели](#указатели)
  - [Указатель на структуру](#указатель-на-структуру)
  - [Почему это важно?](#почему-это-важно)
    - [Иногда имеет большое значение, передать структуру или указатель](#иногда-имеет-большое-значение-передать-структуру-или-указатель)
  - [Чтобы запомнить:](#чтобы-запомнить)
  - [Указатели на срезы](#указатели-на-срезы)
  - [Указатель на ресивер в методе](#указатель-на-ресивер-в-методе)
  - [Разница в определении функций с указателем и без](#разница-в-определении-функций-с-указателем-и-без)
- [nil](#nil)
- [Тесты](#тесты)
- [Переменная интерфейсного типа](#переменная-интерфейсного-типа)
  - [Что внутри переменной интерфейсного типа?](#что-внутри-переменной-интерфейсного-типа)
- [Прикольные фишки](#прикольные-фишки)
  - [Вывод аргумента и значения](#вывод-аргумента-и-значения)
  - [Вывод полей пременной для стуркур и map](#вывод-полей-пременной-для-стуркур-и-map)
  - [Обработка ошибок](#обработка-ошибок)
  - [type assertion](#type-assertion)
  - [Цикл range](#цикл-range)
  - [Метки](#метки)
  - [Опциональные параметры функции](#опциональные-параметры-функции)
  - [Переменное число параметров функции](#переменное-число-параметров-функции)
  - [Замыкания в defer](#замыкания-в-defer)
  - [Указатель на указатель](#указатель-на-указатель)
  - [Кудрявые скобки без оператора](#кудрявые-скобки-без-оператора)
  - [iota](#iota)
  - [Задаём свой формат вывода в Print и аналогах](#задам-свой-формат-вывода-в-print-и-аналогах)
  - [Alias типа](#alias-типа)
  - [Измерение времени кода](#измерение-времени-кода)
- [Без паники! (нет)](#без-паники-нет)
  - [Что такое паника?](#что-такое-паника)
  - [Когда возникает паника?](#когда-возникает-паника)
    - [Go вызывает панику автоматически при:](#go-вызывает-панику-автоматически-при)
  - [recover](#recover)
- [Gorutines](#gorutines)
  - [wg.Wait(), wg.Go() и errgroup.Group()](#wgwait-wggo-и-errgroupgroup)
  - [errgroup.Group](#errgroupgroup)
- [Каналы](#каналы)
    - [Небуфферизованный канал (по умолчанию)](#небуфферизованный-канал-по-умолчанию)
    - [Буфферизованный канал](#буфферизованный-канал)
    - [Однонаправленные каналы](#однонаправленные-каналы)
  - [Select](#select)
    - [Твой компьютер в безопасности](#твой-компьютер-в-безопасности)
    - [Что значит "может привести к утечкам ресурсов"?](#что-значит-может-привести-к-утечкам-ресурсов)
    - [Что стоит делать как программисту:](#что-стоит-делать-как-программисту)
  - [nil-каналы](#nil-каналы)
  - [Резюме](#резюме)
  - [Как завершать горутины красиво?](#как-завершать-горутины-красиво)
- [Generic](#generic)
    - [Что такое generic в Go](#что-такое-generic-в-go)
  - [Пример pipeline](#пример-pipeline)
  - [Важный момент про закрытые каналы](#важный-момент-про-закрытые-каналы)
- [Для примера](#для-примера)
  - [Номер один](#номер-один)
- [Mutex](#mutex)
    - [Зачем он нужен?](#зачем-он-нужен)
    - [Как работает mu.Lock()](#как-работает-mulock)
  - [worker](#worker)
  - [Что такое worker](#что-такое-worker)
- [Мок](#мок)
- [Atomic](#atomic)
  - [Что делает пакет sync/atomic](#что-делает-пакет-syncatomic)
  - [Когда использовать](#когда-использовать)
- [Работа с зависимостями](#работа-с-зависимостями)
- [Init](#init)

*Примечание*. Иногда я буду приводить код вне функции main(), хотя он должен быть в ней.

В таких случаях я буду стараться писать // main

```go
// main

fmt.Println("Hello world!")
```

Для запуска нужно будет скопировать вот часть кода после комментария в функцию main()

и написать название пакета и импорты

# Базовые вещи

## Присваивание

Присваивать значения можно двумя способами.

Второй способ часто используется когда нужно создать переменную прямо на месте

```go
package main

import (
	"fmt"
)

func testingScope() {
	var count int
	for count = 3; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println("Here we go")
	fmt.Printf("count  = %v", count) // Есть доступ сюда
}

func testingScope1() {
	// var count int
	for count := 3; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println("Here we go")
	// fmt.Printf("count  = %v", count)
	// ./example.go:24:28: undefined: count
}

func testingScope2() {
	count := 3
	for count > 0 {
		count--
		fmt.Println(count)
	}

	fmt.Println("Here we go")
	fmt.Printf("count  = %v", count) // Есть доступ сюда
}

func main() {
	testingScope()
}
```

## Как узнать тип переменной?

```go
var x = 42
fmt.Printf("Тип переменной: %T\n", x)

/// Или
import "reflect"

fmt.Println(reflect.TypeOf(x)) // int
```

## Циклы

С предусловием

```go
func main() {
	var count = 0
	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	}
}
```

```go
func main() {
	var count = 0
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
}
```

С коротким объявлением

```go
func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
}
```

Ещё пара вариантов удобного использования короткого присваивания

```go
func hey() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}
```

```go
func lalaley() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("SOme Galactic")
	}
}

```

## Switch

```go
// перебираем возможные варианты для переменной  ID
switch ID {
	// проверяем одно значение
	case "Apple":
	    fmt.Println("Введите свой логин и пароль")
	// проверяем второе значение
	case "Google":
	    fmt.Println("Ваша операционная система не поддерживается")
	// если ничего нужного не нашлось
	default:
	    fmt.Println("Ошибка ввода")
}
```

## Go mod init

Что делает команда

```go
go mod init nameOfProject
```

Команда `go mod init` используется в языке программирования Go для инициализации нового модуля. Модуль в Go — это набор пакетов, которые могут быть собраны и версионированы вместе.

Когда вы выполняете команду `go mod init <module-name>`, она создает файл `go.mod` в текущем каталоге. Этот файл содержит информацию о модуле, включая его имя и зависимости. Вот основные функции этой команды:

1. **Создание файла go.mod**: Файл `go.mod` будет содержать имя вашего модуля и информацию о версиях зависимостей, которые ваш модуль использует.
2. **Управление зависимостями**: После инициализации модуля вы можете добавлять зависимости, и Go будет автоматически управлять их версиями.
3. **Упрощение сборки**: Модули позволяют упростить процесс сборки и управления зависимостями, что делает разработку более организованной.

Таким образом, `go mod init` — это первый шаг в создании модуля Go, который помогает управлять кодом и его зависимостями.

## return

Важный момент

```go
return // просто выйти из функции, не возвращать ничего
```

```go
return nil // вернуть nil
```

Разница есть!

## Массивы

Цикл по массиву создаём копию

```go

a := [1000]int{}
for _, v := range a { /* копия a */ }
```

Лучше

```go

a := [1000]int{}
for _, v := range a[:] { /* копия a */ }
```



# Считывание значений из stdin

```go
package main

import "fmt"

func main() {
	var x int64
	var y float32

	fmt.Scanf("%v %v", &x, &y)
	//fmt.Scanf("%d %f", &x, &y)

	fmt.Printf("x = %v, type %T; y = %v\n, type %T", x, x, y, y)
}
```

Есть ещё сканер Scanf, который позволяет задать более гибкие условия

```go
var a, b int
_, err := fmt.Scanf("%d,%d", &a, &b)
// Ввод: "12,34"
```

```go
var host string
var port int
fmt.Scanf("host=%s port=%d", &host, &port)
// Ввод: "host=localhost port=5432"
```



### fmt.Scanf — чтение по заданному шаблону

fmt.Scanf — это специализированная функция, которая читает текст из стандартного ввода (os.Stdin) и пытается сопоставить его с заданным шаблоном (строкой формата). Она работает очень похоже на scanf в языке C.

По сути, fmt.Scanf(format, a...) — это просто сокращенная запись для fmt.Fscanf(os.Stdin, format, a...).

### Ключевые способности

1. **Форматный ввод**: Это главная особенность. Вы задаете точный шаблон, которому должен соответствовать ввод. Это позволяет читать данные сложной структуры, а не только разделенные пробелами.
2. **Привязка к `os.Stdin`**: Функция жестко привязана к чтению с консоли, что делает ее удобной для быстрых интерактивных утилит.
3. **Автоматическое преобразование типов**: Как и другие функции семейства `Scan`, она автоматически преобразует считанные строки в нужные типы данных (например, `int`, `float64`, `string`) и записывает их по указателям.

### Когда использовать

Используйте `fmt.Scanf`, когда вы ожидаете от пользователя ввод в **строго определённом формате**.

- **Пример 1**: Вы просите пользователя ввести дату в формате `ДД-ММ-ГГГГ`.
- **Пример 2**: Вам нужно считать координаты точки, записанные как `(x, y)`.
- **Пример 3**: Для задач на спортивном программировании, где формат ввода четко оговорен.

```go
package main

import (
    "fmt"
)

func main() {
    var day, month, year int

    fmt.Println("Введите дату в формате ДД-ММ-ГГГГ:")
    // Мы ожидаем два числа, разделенные дефисом
    _, err := fmt.Scanf("%d-%d-%d", &day, &month, &year)
    if err != nil {
        fmt.Println("Ошибка: формат ввода не соответствует ДД-ММ-ГГГГ.", err)
        return
    }

    fmt.Printf("Вы ввели дату: День=%d, Месяц=%d, Год=%d\n", day, month, year)
}
```

### fmt.Fscan* — для простого и типизированного ввода

Функции семейства fmt.Fscan (включая fmt.Fscan, fmt.Fscanln, fmt.Fscanf) идеально подходят для чтения данных, когда вы заранее знаете их формат и тип. Они работают аналогично scanf в языке C, считывая из io.Reader (например, os.Stdin) значения, разделённые пробелами, и записывая их в переменные по указателю.

fmt.Fscan(r io.Reader, a ...interface{}): Читает значения, разделённые пробелами, до тех пор, пока не заполнит все переданные переменные или не встретит ошибку. Переводы строк воспринимаются как пробелы.

fmt.Fscanln(r io.Reader, a ...interface{}): Делает то же самое, но прекращает чтение после первого же перевода строки, если он встретится после считанных аргументов.

fmt.Fscanf(r io.Reader, format string, a ...interface{}): Читает данные в соответствии с заданной строкой формата.

```go
func main() {
    var name string
    var age int

    fmt.Println("Введите ваше имя и возраст через пробел:")
    // Передаём os.Stdin как источник данных
    _, err := fmt.Fscan(os.Stdin, &name, &age)
    if err != nil {
        fmt.Println("Ошибка ввода:", err)
        return
    }

    fmt.Printf("Привет, %s! Тебе %d лет.\\n", name, age)
}
```

**Ключевые особенности:**

- **Удобство:** Очень просто читать числа, строки и другие базовые типы без ручного преобразования.
- **Типизация:** Данные сразу записываются в переменные нужного типа.
- **Обработка ошибок:** **Критически важно** проверять возвращаемую ошибку. Если пользователь введёт текст вместо числа, функция вернёт ошибку, и переменная останется без изменений.

**Когда использовать:** Для простых консольных утилит, где нужно быстро считать несколько значений (например, два числа для калькулятора).

### bufio.Scanner — простой API для строк и токенов

bufio.Scanner — это, пожалуй, самый удобный и идиоматичный способ для построчного чтения ввода в Go. Он скрывает всю сложность буферизации и поиска разделителей, предоставляя очень простой API.

По умолчанию Scanner разбивает ввод на строки, но его можно настроить для разбиения по словам (bufio.ScanWords) или даже по собственным правилам.

**Ключевые особенности:**

- **Простота:** Очень лаконичный и понятный API.
- **Надёжность:** Хорошо справляется с различными окончаниями строк (`\n`, `\r\n`).
- **Ограничение:** По умолчанию максимальный размер одного токена (например, строки) ограничен 64 КБ. Это можно изменить с помощью метода `scanner.Buffer()`, но об этом нужно помнить.

**Когда использовать:** В большинстве случаев, когда вам нужно читать стандартный ввод **построчно**. Это предпочтительный метод для этой задачи.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Создаём новый Scanner для стандартного ввода
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Введите несколько строк текста (для завершения введите 'exit'):")

    // Цикл будет выполняться для каждой новой строки
    for scanner.Scan() {
        line := scanner.Text() // Получаем текст текущей строки
        if line == "exit" {
            break
        }
        fmt.Println("Вы ввели:", line)
    }

    // Проверяем на ошибки, которые могли возникнуть во время сканирования
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
    }
}
```

### bufio.Reader — гибкость и контроль

bufio.Reader предоставляет буферизованное чтение из io.Reader. Это означает, что он считывает данные из источника "порциями" (чанками) в свой внутренний буфер, а затем отдает их вашей программе. Это значительно эффективнее, чем читать по одному байту за раз.

Основной метод — ReadString('\n') или ReadBytes('\n') — позволяет читать данные до определённого разделителя (чаще всего до символа новой строки \n).

**Ключевые особенности:**

- **Гибкость:** Можно читать данные до любого символа-разделителя, а не только до пробела.
- **Производительность:** Буферизация снижает количество системных вызовов, что полезно при работе с файлами или сетевыми соединениями.
- **Контроль:** Даёт больше контроля над процессом чтения, чем `fmt.Fscan`.

**Когда использовать:** Когда нужно читать ввод построчно или по другим разделителям, особенно если строки могут быть длинными.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Создаём новый Reader для стандартного ввода
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Введите ваше полное имя:")

    // Читаем строку до символа новой строки
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Ошибка чтения:", err)
        return
    }

    // Удаляем символ новой строки из конца строки
    name := strings.TrimSpace(input)
    fmt.Printf("Здравствуйте, %s!\n", name)
}
```

### io.Reader.Read — низкоуровневый контроль

Это самый фундаментальный и низкоуровневый способ чтения. Метод Read интерфейса io.Reader (который реализует os.Stdin) просто пытается прочитать данные в предоставленный срез байт ([]byte).

Вы получаете полный контроль, но и всю ответственность на себя.

**Ключевые особенности:**

- **Максимальный контроль:** Вы управляете размером буфера и сами обрабатываете каждый прочитанный байт.
- **Сложность:** Необходимо вручную обрабатывать ситуации, когда данные приходят "кусками" (частичное чтение). Например, вы запросили 1024 байта, а система вернула только 100.
- **Эффективность:** Может быть очень эффективным, если вы точно знаете, как управлять памятью и буферами для вашей конкретной задачи (например, при написании сетевого протокола).

**Когда использовать:** При работе с бинарными данными, реализации собственных протоколов или в ситуациях, где требуется максимальная производительность и минимальные накладные расходы. Для чтения простого текстового ввода от пользователя этот метод избыточен.

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // Создаём буфер (срез байт) для чтения данных
    buffer := make([]byte, 1024)
    fmt.Println("Введите что-нибудь:")

    // Читаем данные из os.Stdin в наш буфер
    n, err := os.Stdin.Read(buffer)
    if err != nil && err != io.EOF {
        fmt.Println("Ошибка чтения:", err)
        return
    }

    // Выводим только те байты, которые были реально прочитаны (от 0 до n)
    fmt.Printf("Прочитано %d байт: %s\n", n, string(buffer[:n]))
}
```

# Импорты

Существует несколько форм оператора import

- Обычный импорт
```
import "fmt"

fmt.Println()
```

- Импорт с псевдонимом

```
import f "fmt"


f.Println() // не fmt.Println()
```

- Импорт с точкой (не рекоменудется)

```
import . "fmt"

Println("Something") // без указания пакета
```

- Пустой импорт (только для инициализации пакета (функция init), не для прямого использования)

```
import _ package
```

# Директории и переменные окружения

```bash

go env GOROOT GOPATH GOMODCACHE GOBIN
```

- GOROOT: путь к дистрибутиву.
- GOPATH: корень рабочего пространства.
- GOMODCACHE: кеш модулей (обычно $GOPATH/pkg/mod).
- GOBIN: куда устанавливаются бинарники при go install.

## Платформы и кросс-компилляция

Можно компиллировать программы для разных систем

```bash

GOOS=linux GOARCH=arm64 go build ./cmd/app
```

```bash
GOOS=linux GOARCH=amd64 go build -o myprog-linux-amd64 ./...
```

```bash
GOOS=windows GOARCH=amd64 go build -o myprog-windows-amd64.exe ./...
```

GOTOOLCHAIN в Go (Golang) — это переменная среды, которая указывает на инструменты компиляции Go. Она позволяет управлять инструментами разработки и сборки, такие как компиляторы, линкеры и другие утилиты.

# Распаковка

```go
func PrintAll(words ...string) {}

PrintAll("a", "b", "c") // можно так

// или так
list := []string{"a", "b", "c"}
PrintAll(list...) // распаковка
```

```go
...prefix // в аргументах — позволяет передавать любое количество аргументов

prefix... // в вызове — распаковывает слайс в отдельные аргументы.
```

```go

```

# defer

Возможность отложить какое-то действие. Оно выполнится после окончания функции, после return, действие будет выполнено даже если возникнет паника.

Удобно, для очистки ресурсов

```go
func example() {
	defer fmt.Println("Первый")
	defer fmt.Println("Второй")
	defer fmt.Println("Третий")
}

// Третий
// Второй
// Первый

```

Здесь defer выполняет различные действия в зависимости от возвращаемой функцией ошибки

```go

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// здесь можно выполнить еще ряд операций вставки, используя tx
	return nil
}

```

## defer не выполняетя при os.Exit()

# ООП в golang

Наследование, инкапсуляция и полиморфизм — это три основных принципа объектно-ориентированного программирования (ООП). Давайте рассмотрим каждый из них и приведем примеры на языке Go (Golang).

## 1. Наследование

Наследование позволяет создавать новые структуры (или классы), которые наследуют свойства и методы существующих структур. В Go наследование реализуется через встраивание.

```go
package main

import "fmt"

// Определяем базовую структуру
type Animal struct {
	Name string
}

// Метод для базовой структуры
func (a Animal) Speak() {
	fmt.Println("Animal speaks")
}

// Определяем структуру, которая наследует Animal
type Dog struct {
	Animal // Встраивание
	Breed  string
}

// Переопределяем метод Speak
func (d Dog) Speak() {
	fmt.Println("Woof! My name is", d.Name)
}

func main() {
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}
	dog.Speak() // Вывод: Woof! My name is Buddy
}
```

## 2. Инкапсуляция

Инкапсуляция — это механизм, который ограничивает доступ к внутренним данным и методам объекта. В Go это достигается с помощью экспортируемых и неэкспортируемых идентификаторов (с заглавной и строчной буквы соответственно).

```go
package main

import "fmt"

// Структура с инкапсуляцией
type Person struct {
	name string // неэкспортируемое поле
	Age  int    // экспортируемое поле
}

// Метод для установки имени
func (p *Person) SetName(name string) {
	p.name = name
}

// Метод для получения имени
func (p *Person) GetName() string {
	return p.name
}

func main() {
	person := Person{Age: 30}
	person.SetName("Alice")
	fmt.Println("Name:", person.GetName()) // Вывод: Name: Alice
}
```

## 3. Полиморфизм

Полиморфизм позволяет использовать один интерфейс для работы с различными типами. В Go полиморфизм реализуется через интерфейсы.

```go
package main

import "fmt"

// Определяем интерфейс
type Speaker interface {
	Speak()
}

// Структура Cat
type Cat struct {
	Name string
}

// Реализация метода Speak для Cat
func (c Cat) Speak() {
	fmt.Println("Meow! My name is", c.Name)
}

// Структура Dog
type Dog struct {
	Name string
}

// Реализация метода Speak для Dog
func (d Dog) Speak() {
	fmt.Println("Woof! My name is", d.Name)
}

// Функция, принимающая интерфейс Speaker
func makeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	cat := Cat{Name: "Whiskers"}
	dog := Dog{Name: "Buddy"}

	makeItSpeak(cat) // Вывод: Meow! My name is Whiskers
	makeItSpeak(dog) // Вывод: Woof! My name is Buddy
}
```

В этих примерах мы продемонстрировали основные принципы ООП в Go: наследование через встраивание, инкапсуляцию с использованием экспортируемых и неэкспортируемых полей, а также полиморфизм через интерфейсы.

Структурирование кода

Например, есть такой код

```go
package main

import (
	"fmt"
)

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type report struct {
	sol int
	temperature
	location
}

// Объявим метод для подсчёта средней (арифметической) температуры
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// Функция обёртка
// Не обязательна в данном случае
// Нужна только, если бы мы хотели добавить какую-то дополнительную логику
func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	t := temperature{high: -1.0, low: -78.0}

	report := report{sol: 15, temperature: t}
	fmt.Printf("average %vº C\n", report.temperature.average())
	fmt.Printf("average %vº C\n", report.average())
}

```

Удобно же, правда? Но можно проще)

> Функция обёртка нужна только если мы хотим сделать дополнительную логику. Можно и без неё
>

## Встраивания (Эмбединги)

```go
package main

import (
	"fmt"
)

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type report struct {
	sol int
	temperature
	location
}

// Объявим метод для подсчёта средней (арифметической) температуры
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	report := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}

	// Можно вызвать как
	fmt.Printf("average %vº C\n", report.average())

	// Так же можно вызвать
	fmt.Printf("average %vº C\n", report.temperature.average())
}

```

Ещё один пример встраивания в Golang

```go
package main

import "fmt"

// создадим alias
type sol int

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type report struct {
	sol
	location    location
	temperature temperature
}

// создадим метод структуры sol
func (s sol) days(s2 sol) sol {
	days := s2 - s
	if days < 0 {
		days = -days
	}
	return days
}

func main() {
	report := report{sol: 15}

	duration := sol(1446)

	fmt.Println(report.sol.days(duration))
	fmt.Println(report.days(duration))
}

```

> **Важный нюанс**
>

```go
// Так работать будет
type report struct {
	sol
	location    location
	temperature temperature
}

// А так уже нет
type report struct {
	sol sol
	location    location
	temperature temperature
}
```

Без имени поля, только sol — это называется **встраивание** (**embedding**). Это как будто report «наследует» все методы и поля sol напрямую. То есть report становится как бы солом, и ты можешь писать:

```go
report.days(duration)
```

Потому что report сам знает про метод days, благодаря встраиванию.

С указанием имени поля — это обычное поле. Тут sol — это просто переменная внутри структуры, а не встраивание.  В этом случае report уже не знает про метод days, потому что sol внутри как коробочка, и метод days нужно вызывать через неё:

```go
report.sol.days(duration) // только так

//уже не работает, потому что report сам этот метод не получил
report.days(duration) // не работает
```

Как написано | Что происходит | Как вызывать метод
 | Встраивание. Методы доступны у report. | report.days(...)
| Просто поле. Методы не доступны у report. | report.sol.days(...)

| Как написано | Что происходит | Как вызвать метод |
| --- | --- | --- |
| sol | Встраивание. Доступны все методы report | report.days(…) |
| sol sol | Просто поле. Методы report не доступны | report.sol.days(…) |

### Как понять когда писать с именем, когда без:

- Если ты хочешь, чтобы структура сразу «умела» всё, что умеет вложенный тип — **пиши без имени**(embedding).
- Если это просто данные и ты не хочешь напрямую связывать типы — **пиши с именем** (обычное поле).

## Ещё примеры встраивания

Без встраивания

```go
package main

import "fmt"

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Двигатель запущен")
}

type Car struct {
	engine Engine  // обычное поле
}

func main() {
	myCar := Car{engine: Engine{}}
	myCar.engine.Start()  // доступ через поле
}

```

Со встраиванием

```go
package main

import "fmt"

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Двигатель запущен")
}

type Car struct {
	engine Engine  // обычное поле
}

func main() {
	myCar := Car{engine: Engine{}}
	myCar.engine.Start()  // доступ через поле
}

```

# Интерфейсы в Golang

Интерфейс – это набор методов. Если у типа есть все необходимые методы, автоматически считается, что он реализует этот интерфейс.

Например, есть интерфейс “Пульт дистанционного управления”, у которого два метода: включить и выключить.

Этот пульт подойдёт для любого устройства (структура), у которой есть методы включить и выключить (Телевизор, кондиционер, вибратор)

Пример

```go
package main

import "fmt"

// Define an interface
type Animal interface {
	Speak() string
}

// Dog type
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat type
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Use the interface
func makeItSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	d := Dog{}
	c := Cat{}

	makeItSpeak(d) // Output: Woof!
	makeItSpeak(c) // Output: Meow!
}
```

И собака, и кошка реализуют функцию Speak(), поэтому они автоматически реализуют интерфейс Animal.

## Еще одно объяснение интефейсов

Типы данных фокусируются на том, какие данные в них хранятся (int, string, float64 и тд)

Методы выражают поведение типа.

Интерфейсы как раз фокусируются на поведении типов, а не на том, какие данные хранятся.

```go
package main

import (
	"fmt"
	"strings"
)

// t может хранить любой объект, если у него есть
// метод talk(), который возвращает строку.
var t interface {
	talk() string
}

// Это структура martian. У неё есть метод talk(), который возвращает
// строку "nack nack".
// Она удовлетворяет интерфейсу, потому что у неё есть
// нужный метод.
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

// тоже ужовлетворяет интерфейсу
// тоже есть методом talk() string
type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {
		var t interface {
			talk() string
			}
			t = martian{}
			fmt.Println(t.talk())
			t = laser(3)
		fmt.Println(t.talk())
}

```

***Примечание***

Обычно называют интерфейсы с суффикосом -er.

talker – всё, у чего есть метод talk

```go
type talker interface {
		talk() string
}

// Например, это можно использовать
// вот так в дальнешем
func shout(t talker) {
		louder := strings.ToUpper(t.talk())
		fmt.Println(louder)
}
```

## Использование интерфейсов

В Go внутри пакета fmt (который отвечает за печать в консоль)

есть специальный интерфейс

```go
type Stringer interface {
    String() string
}
```

Это значит: 👉 если твой тип (например, location) имеет метод String() string, то fmt.Println() будет вызывать этот метод, чтобы понять как красиво напечатать твой объект.

Если без метода String()

```go
type location struct {
    lat, long float64
}

curiosity := location{-4.5895, 137.4417}
fmt.Println(curiosity) // {-4.5895 137.4417}
```

С методом

```go
func (l location) String() string {
    return fmt.Sprintf("%v, %v", l.lat, l.long)
}

fmt.Println(curiosity) // -4.5895, 137.4417
```

fmt.Println() проверяет:
**а реализует ли объект интерфейс Stringer?**

- Если да — вызывает твой метод String()
- Если нет — печатает как есть (struct со скобками и полями).

То есть можно самому определить, как будет выглядеть Println твоего объекта

### Как можно использовать интерфейсы в тестах?

Структуры и интерфейсы позволяют подменять реальные реализации на фейковые в тестах

Например, был вот такой код

```go
var DB *sql.DB // глобальная переменная

func GetUserName(id int) (string, error) {
	var name string
	err := DB.QueryRow("SELECT name FROM users WHERE id = $1", id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

```

👉 Минусы:

- в тесте `DB` всё равно настоящий (или нужно мутить глобальный мок);
- тесная привязка к Postgres.

С использованием интерфейсов

```go
// Объявим интерфейс
type Storage interface {
	GetUserName(id int) (string, error)
}

// И функцию, которая использует этот интерфейс
func PrintUserName(storage Storage, id int) {
	name, err := storage.GetUserName(id)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("User:", name)
}

```

Реализация для реальной базы данных

```go
type SQLStorage struct {
	DB *sql.DB
}

func (s *SQLStorage) GetUserName(id int) (string, error) {
	var name string
	err := s.DB.QueryRow("SELECT name FROM users WHERE id = $1", id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

```

Реализация для тестов

```go
// Структура-мок
type FakeStorage struct{}

// Функция-мок
func (f *FakeStorage) GetUserName(id int) (string, error) {
	return "TestUser", nil
}

```

Как в итоге использовать?

```go
func main() {
	// 🔹 боевой режим
	realStorage := &SQLStorage{DB: realDB}
	PrintUserName(realStorage, 1)

	// 🔹 тестовый режим
	fakeStorage := &FakeStorage{}
	PrintUserName(fakeStorage, 42) // всегда вернёт "TestUser"
}
```

# Паттерны проектирования

## Фабрика

Фабрика – это функция, которая создаёт и возвращает объект

```go
// Структура, экземпляры которой нужно создавать
type fakeEntry struct {
	name  string
	isDir bool
}

// Фабрики
func fakeFile(name string) os.DirEntry {
	return &fakeEntry{name: name, isDir: false}
}

func fakeDir(name string) os.DirEntry {
	return &fakeEntry{name: name, isDir: true}
}

// Без фабрики
entries: []os.DirEntry{
	&fakeEntry{name: "f1", isDir: false},
	&fakeEntry{name: "f2", isDir: false},
	&fakeEntry{name: "d1", isDir: true},
}

// С фабриков
entries: []os.DirEntry{
	fakeFile("f1"),
	fakeFile("f2"),
	fakeDir("d1"),
}
```

Вот ещё пример

```go
package main

import "fmt"

// Структура: Человек
type Person struct {
	name string
	age  int
}

// Фабричная функция для создания взрослого человека
func NewAdult(name string) Person {
	return Person{name: name, age: 30}
}

// Фабричная функция для создания ребёнка
func NewChild(name string) Person {
	return Person{name: name, age: 5}
}

func main() {
	// Используем фабрики вместо ручного создания
	alice := NewAdult("Alice")
	bobby := NewChild("Bobby")

	fmt.Println(alice) // Вывод: {Alice 30}
	fmt.Println(bobby) // Вывод: {Bobby 5}
}

```

## Адаптер

Если нужно адаптировать старый код под новый интерфейс

```go
package main

import "fmt"

// Представим: у нас есть игра, где все персонажи должны уметь атаковать
type Fighter interface {
	Attack() string
}

// У нас есть рыцарь - он умеет атаковать
type Knight struct {
	name string
}

func (k *Knight) Attack() string {
	return fmt.Sprintf("%s атакует мечом!", k.name)
}

// Но есть старый код робота из другой игры
// У робота другой метод - он стреляет, а не атакует
type OldRobot struct {
	model string
}

func (r *OldRobot) Shoot() string {
	return fmt.Sprintf("Робот %s стреляет лазером!", r.model)
}

// Адаптер! Он "переводит" робота в понятный для игры формат
type RobotAdapter struct {
	robot *OldRobot
}

// Адаптер делает вид, что робот умеет атаковать (как все остальные)
func (adapter *RobotAdapter) Attack() string {
	// Внутри он просто вызывает метод стрельбы робота
	return adapter.robot.Shoot()
}

// Функция боя - она работает с любым Fighter
func StartBattle(fighter Fighter) {
	fmt.Println(fighter.Attack())
}

func main() {
	// Обычный рыцарь
	knight := &Knight{name: "Артур"}

	// Старый робот (не подходит для игры напрямую)
	oldRobot := &OldRobot{model: "T-800"}

	// Адаптер для робота
	robotAdapter := &RobotAdapter{robot: oldRobot}

	fmt.Println("=== БОЙ НАЧИНАЕТСЯ ===")

	// Рыцарь может сражаться сразу
	StartBattle(knight)

	// Робот может сражаться только через адаптер
	StartBattle(robotAdapter)

	fmt.Println("\n=== Все бойцы в одном списке ===")

	// Теперь можем всех добавить в один список!
	fighters := []Fighter{knight, robotAdapter}

	for _, fighter := range fighters {
		StartBattle(fighter)
	}
}
```

- Есть интерфейс `AudioPlayer` с методом `Play(file string)`.
- Есть старый `CassettePlayer` с методом `PlayTape(filename string)`.
- Сделай адаптер.

```go
package main

import "fmt"

type AudioPlayer interface {
	Play(file string)
}

// Старая структура
type CassettePlayer struct {
	year string
}

// Метод старой структуры
func (c *CassettePlayer) PlayTape(filename string) {
	fmt.Println("Playing ", filename)
}

type Adapter struct {
	old *CassettePlayer
}

func (a *Adapter) Play(file string) {
	newString := "New " + file
	a.old.PlayTape(newString)
}

func main() {
	oldThing := &CassettePlayer{year: "2000"}
	newThing := &Adapter{oldThing}
	newThing.Play("My favorite music")
	fmt.Println(newThing.old.year)
}

```

Ещё пример (из реального проекта)

```go
package main

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
)

// Создали интерфейс зависимости
type UserManager interface {
	ListMembers(ctx context.Context, key string) ([]string, error)
}

// Наша структура зависит от интерфейса
type AdminService struct {
	UserManager UserManager
}

// Фабрика AdminService с зависимостью от
// интерфейса
func NewAdminService(userManager UserManager) *AdminService {
	return &AdminService{
		UserManager: userManager,
	}
}

// Тот самый адаптер
//
// # Оборачиваем в свою структуру реальную структуру
//
// Реализует интерфейс UserManager
type RedisUserManager struct {
	RedisClient *redis.Client
}

// Реализуем интерфейс
//
// Фактическая реализация интерфейса
func (r *RedisUserManager) ListMembers(ctx context.Context, key string) ([]string, error) {
	return r.RedisClient.SMembers(ctx, key).Result()
}

// Фабрика UserManager
func NewRedisUserManager(client *redis.Client) *RedisUserManager {
	return &RedisUserManager{RedisClient: client}
}

// Моковая реализация
type MockUserManager struct {
	mock.Mock
}

// Фабрика MockUserManager
func NewMockUserManager() *MockUserManager {
	return &MockUserManager{}
}

// Создаём мок
//
// # Реализуем интерфейс
//
// Моковая реализация интерфейса
func (m *MockUserManager) ListMembers(ctx context.Context, key string) ([]string, error) {
	args := m.Called(ctx, key)
	return args.Get(0).([]string), args.Error(1)
}

func main() {
	// Для production
	userManager := NewRedisUserManager(client * redis.Client)
	AdminService := NewAdminService(userManager)

	// Для тестов
	mockUserManager := NewMockUserManager()
	AdminService = NewAdminService(mockUserManager)
}

```

То есть адаптер просто оборачивает в нашу структуру реальную реализацию. А дальше для своей структуры мы можем определять любые методы для реализации нашего интерфейса

### Когда применять:

Паттерн Адаптер особенно полезен при интеграции с внешними API, миграции между версиями библиотек, или когда нужно унифицировать работу с разными сервисами под общим интерфейсом.

Или

- У тебя есть **готовый старый код** (робот), который работает, но интерфейс не подходит
- Не хочется **переписывать** весь старый код
- Нужно чтобы **новый и старый код работали вместе**

Адаптер решает это - он как "мостик" между старым и новым кодом!

Ещё пример

## Dependency Injection (Внедрение зависимостей)

Это когда код получает зависимости снаружи (а не когда зависимости зашиты внутри кода)

Без DI (жёсткая зависимость)

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRepo struct{}

func (r *UserRepo) GetUser() string {
	return "Alice"
}

// Вот здесь проблема. UserHandler обязан
// использовать структуру UserRepo
// Использовать другой репозиторий нельзя
type UserHandler struct {
	repo *UserRepo
}

// И здесь нам тоже уже приходится использовать
// именно UserRepo
func NewUserHandler() *UserHandler {
	return &UserHandler{repo: &UserRepo{}}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	c.String(http.StatusOK, h.repo.GetUser())
}

func main() {
	r := gin.Default()

	handler := NewUserHandler()
	r.GET("/user", handler.GetUser)

	r.Run(":8080")
}

```

C DI

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Интерфейс репозитория
type UserRepo interface {
	GetUser() string
}

// Реальная БД
type DbUserRepo struct{}

// Реализуем интерфейс
func (r *DbUserRepo) GetUser() string {
	return "Alice from DB"
}

// Фейковая БД для тестов
type MockUserRepo struct{}

// Реализуем интерфейс
func (r *MockUserRepo) GetUser() string {
	return "Alice from Mock"
}

// Зависимость от интерфейса
// То есть можем использовать всё, что реализует интерфейс
type UserHandler struct {
	repo UserRepo
}

// Тут зависимость передаётся снаружи
// В качетсве аргумента можно передать всё, что реализует
// интерфейс
func NewUserHandler(repo UserRepo) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	c.String(http.StatusOK, h.repo.GetUser())
}

// В production
func main() {
	r := gin.Default()

	// В реальном коде передаём реальную БД
	handler := NewUserHandler(&DbUserRepo{})
	r.GET("/user", handler.GetUser)

	r.Run(":8080")
}

// В тестах
func main() {
	r := gin.Default()

	// А тут, например, мы передаём мок (для тестов)
	handler := NewUserHandler(&MockUserRepo{})
	r.GET("/user", handler.GetUser)

	r.Run(":8080")
}

```

Ещё пример

Без DI

```go
package main

import "fmt"

type UserRepo struct{}

func (r *UserRepo) GetUser() string {
	return "Alice"
}

// Жёсткая зависимость
type UserService struct {
	repo *UserRepo
}

// Приходится использовать именно UserRepo
func NewUserService() *UserService {
	return &UserService{repo: &UserRepo{}}
}

func (s *UserService) PrintUser() {
	fmt.Println(s.repo.GetUser())
}

func main() {
	service := NewUserService()
	service.PrintUser()
}

```

С DI

```go
package main

import "fmt"

type UserRepo interface {
	GetUser() string
}

// Зависимость от интерфеса
// То есть от набора методов
type UserService struct {
	repo UserRepo
}

// Реальная база данных
type DbUserRepo struct{}

// Реализуем интерфейс
func (r *DbUserRepo) GetUser() string {
	return "Alice from DB"
}

// Мок для тестов
type MockUserRepo struct{}

// Реализуем интерфейс
func (r *MockUserRepo) GetUser() string {
	return "Alice from Mock"
}

// Зависимость передаётся "снаружи"
// Подставляем всё, что реализует интерфейс
func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) PrintUser() {
	fmt.Println(s.repo.GetUser())
}

func main() {
	// В реальном коде используем базу
	service := NewUserService(&DbUserRepo{})
	service.PrintUser()

	// В тестах используем мок
	testService := NewUserService(&MockUserRepo{})
	testService.PrintUser()
}

```

### Алгоритм создания di

1. Создаем интерфейс для зависимости (что зависимость должна уметь?)

```go
// Создаём интерфейс для зависимости (что зависимость должна уметь?)
type UserRepo interface {
    GetUser() string
}

// Создаём одну или несколько реализаций интерфейса (структуру и методы для неё)
type DbUserRepo struct{}
func (r *DbUserRepo) GetUser() string { return "Alice from DB" }

type MockUserRepo struct{}
func (r *MockUserRepo) GetUser() string { return "Alice from Mock" }

// Сохраняем зависимость в структуре, которая её использует (через интерфейс)
type UserHandler struct {
    repo UserRepo // интерфейс, а не конкретная структура!
}

// Делаем фабрику, которая принимает зависимость снаружи
func NewUserHandler(repo UserRepo) *UserHandler {
    return &UserHandler{repo: repo}
}

// Подставляем нужную зависимость в код

// Для реальной работы
handler := NewUserHandler(&DbUserRepo{})

// Для тестов
handler := NewUserHandler(&MockUserRepo{})
```

# Перенаправления методов (перегрузка)

```go
package main

import (
	"fmt"
)

// / Типы данных
type sol int
type celsius float64

// / Структуры
type report struct {
	sol
	location
	temperature
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

/// Функции

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	// To-do: complicated distance calculation
	fmt.Println(l2)
	return 5
}

// "Перегрузим" функцию
// Чтобы избежать коллизий имен
// Здесь мы фактически переадресовали вызов на метод sol.days()
func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {

	t := temperature{
		high: 30.0,
		low:  15.0,
	}

	loc := location{
		lat:  0.0,
		long: 0.0,
	}

	r1 := report{
		1777,
		loc,
		t,
	}

	d := r1.days(1446)
	_ = d

	fmt.Println(d)
}

```

# Замыкания

***Замыкание*** — это функция, которая запоминает переменные из той области, где она была создана, даже если эта область уже завершила работу.

```go
func counter() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

c := counter()

fmt.Println(c())  // 1
fmt.Println(c())  // 2
fmt.Println(c())  // 3
```

Сложный пример

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d, pos(i) = %d, neg(-2i) = %d\n", i, pos(i), neg(-2*i))
	}
}

// i = 0, pos(i) = 0, neg(-2i) = 0
// i = 1, pos(i) = 1, neg(-2i) = -2
// i = 2, pos(i) = 3, neg(-2i) = -6
// i = 3, pos(i) = 6, neg(-2i) = -12
// i = 4, pos(i) = 10, neg(-2i) = -20
// i = 5, pos(i) = 15, neg(-2i) = -30
// i = 6, pos(i) = 21, neg(-2i) = -42
// i = 7, pos(i) = 28, neg(-2i) = -56
// i = 8, pos(i) = 36, neg(-2i) = -72
// i = 9, pos(i) = 45, neg(-2i) = -90
```

Объяснение

| i | pos(i) | Что делает pos | neg(-2*i) | Что делает neg |
| --- | --- | --- | --- | --- |
| 0 | pos(0)=0 | sum=0+0=0 | neg(0)=0 | sum=0+0=0 |
| 1 | pos(1)=1 | sum=0+1=1 | neg(-2)=-2 | sum=0+(-2)=-2 |
| 2 | pos(2)=3 | sum=1+2=3 | neg(-4)=-6 | sum=-2+(-4)=-6 |
| 3 | pos(3)=6 | sum=3+3=6 | neg(-6)=-12 | sum=-6+(-6)=-12 |
| 4 | pos(4)=10 | sum=6+4=10 | neg(-8)=-20 | sum=-12+(-8)=-20 |

Для положительных (для i=2): у нас уже в переменной хранилось значение 1, мы добавили туда 2

Для отрицательных (для i=2): у нас уже было значение -2, мы к нему прибавляем -2*i, то есть -2+(-2*2)=6

# Generic

```go
// Обычные функции

func SumInt(a, b int) int {
    return a + b
}
func SumFloat(a, b float64) float64 {
    return a + b
}

// Обобщённые generic
func Sum[T int | float64](a, b T) T {
    return a + b
}

```

💡 Для чего это нужно:

- Чтобы не дублировать код для разных типов данных.
- Делает код более читаемым и безопасным.
- Упрощает поддержку и уменьшает количество ошибок.

# Resful API

💡 REST API особенности:

- Использует HTTP (методы: GET, POST, PUT, DELETE).
- Без состояния: каждый запрос самодостаточен *
- Ответы обычно в JSON формате.
- Это простой, понятный и удобный способ для связи между программами.

* Самодостаточность запросов

👉 **Сервер не запоминает, кто ты и что ты делал до этого.**

Каждый раз, когда клиент (твое приложение) отправляет запрос, он должен **все важное передать в этом же запросе**:

- токен авторизации,
- параметры,
- данные,
- настройки.

### Почему это круто

- Сервер проще, не нужно хранить "сессии" и "кто с кем общался".
- Любой запрос можно обработать отдельно.
- Легко масштабировать (серверы могут быть разными, запросы идут на любой из них).

✅ **Запомни коротко:**

**Stateless** = сервер не держит в памяти контекст между запросами. Всё, что нужно — клиент передаёт каждый раз

```go
┌────────┬───────────────────────────────┐
│  GET   │ 🔍 Получить данные            │
│        │ (ничего не меняет)            │
│        │ Пример: "Покажи список!"      │
├────────┼───────────────────────────────┤
│  POST  │ 🧾 Создать новый объект        │
│        │ Пример: "Добавь нового!"      │
├────────┼───────────────────────────────┤
│  PUT   │ ♻️ Полностью обновить объект  │
│        │ Пример: "Вот новые данные!"   │
├────────┼───────────────────────────────┤
│ DELETE │ ❌ Удалить объект              │
│        │ Пример: "Удали этот элемент!" │
└────────┴───────────────────────────────┘

```

# Указатели

Бу! Уже боитесь?)

Основные понятия

Амперсанд &

- Возвращает место в памяти компьютера

Астерикс *

- Получает значение по месту в памяти компьютера
- указатель (например при указании типа)

```go
var home *string
```

Пример

```go
// Переопределим Printf для удобства
func Printf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func main() {
	// знак * может встречаться и в объявлении переменных
	var place *string

	var fruit string = "Apple"

	place = &fruit

	// Посмотрим значения переменных
	Printf("fruit = %v", fruit)
	Printf("place = %v", place)

	// Посмотрим типы
	Printf("type of fruit = %T", fruit)
	Printf("type of place = %T", place)

}
```

```go
// Переопределим Printf для удобства

func Printf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func main() {
	// знак * может встречаться и в объявлении переменных
	var home *string

	country := "Canada"

	home = &country

	// Посмотрим теперь
	Printf("Type of home is %T\n", home)  // *sting
	Printf("Type of country %T", country) // string

}
```

Пример, который поясняет всё
Вероятно, можно оставить только его

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func LexPrint(args ...interface{}) {
	// Если аргументов нет — просто выводим пустую строку
	if len(args) == 0 {
		fmt.Fprint(os.Stdout, "\n")
		return
	}

	// Первый аргумент — строка формата
	format, ok := args[0].(string)
	if !ok {
		panic("LexPrint: first argument must be a format string")
	}

	// Остальные — подставляемые значения
	a := args[1:]

	s := fmt.Sprintf(format, a...)
	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}
	fmt.Fprint(os.Stdout, s)
}

func main() {
	// указаель
	var administrator *string

	name1 := "Christopher J. Scolese"

	// теперь указатель указывает на name1
	administrator = &name1
	LexPrint(*administrator)

	name2 := "Charles F. Bolden"

	// теперь указатель указывает на name2
	administrator = &name2
	LexPrint(*administrator)

	// А теперь фокус!
	// Меняем значение, указатель
	// тоже обновляется
	name2 = "Changing1"
	LexPrint(*administrator) // Changing1

	// Дальше больше
	// Меня значение ПО указателю
	*administrator = "Changing2"
	LexPrint(name2) // Changing2

	// fmt.Println()
	LexPrint()

	// Ещё интересная вещь
	dog1 := "Buddy"
	pDog1 := &dog1

	// Можно было просто сделать dog2 := dog1
	dog2 := *pDog1

	// Изменим значение по указателю
	*pDog1 = "BestFriend"

	LexPrint(dog1) // BestFriend
	LexPrint(dog2)
}
```

## Указатель на структуру

```go
package main

import (
	"fmt"
)

type box struct {
	weight int
}

func main() {
	b := box{weight: 10}  // просто структура
	p := &box{weight: 20} // указатель на структуру

	fmt.Printf("b = %v \n", b)
	fmt.Printf("p = %v \n", p)

	fmt.Printf("b.weight = %v \n", b.weight) // 10
	fmt.Printf("p.weight = %v \n", p.weight) // 20

	b.weight = 15 // меняем структуру
	p.weight = 25 // меняем структуру через указатель

	fmt.Printf("b.weight = %v \n", b.weight) // 15
	fmt.Printf("p.weight = %v \n", p.weight) // 25
}
```

Смысл в том, что передача указателя на структуру – потребляет меньше ресурсов.

## Почему это важно?

- **Указатель легче копировать**, чем всю структуру (особенно если структура большая).
- **Можно изменять структуру через указатель**, и изменения сохраняются.
- **Экономия памяти и скорости**, если структур много.

В Go — удобная фишка:

**Если у тебя указатель на структуру, ты всё равно можешь обращаться к полям через `.`**, а не через `*`.

| Без указателя | С указателем |
| --- | --- |
| Копируется весь объект | Передаётся только адрес |
| Изменения внутри функции не видны снаружи | Изменения сохраняются |
| Больше памяти и работы при больших структурах | Эффективно по памяти |

```go
| Без указателя              | С указателем                |
|----------------------------|-----------------------------|
| Копируется весь объект     | Передаётся только адрес     |
| Изменения не сохраняются   | Изменения сохраняются       |
| Больше памяти и операций   | Меньше памяти и операций    |
| Оригинал остаётся нетронут | Оригинал может измениться   |
|----------------------------|-----------------------------|
```

### Иногда имеет большое значение, передать структуру или указатель

```go
package main

import "fmt"

type person struct {
	name string
}

// изменяем имя
func changeName(p person) {
	p.name = "Changed!"
}

// изменяем имя через указатель
func changeNameP(p *person) {
	p.name = "Changed!"
}

func main() {
	p := person{name: "Original"}
	changeName(p)       // передаём КОПИЮ
	fmt.Println(p.name) // ➡️ Выведет: Original
	fmt.Println("Теперь передаём указатель")
	changeNameP(&p)     // передаём УКАЗАТЕЛЬ
	fmt.Println(p.name) // ➡️ Выведет: Changed!
}
```

## Чтобы запомнить:

> Если хочешь менять оригинальные данные — передавай указатель.
>

> Если хочешь оставить оригинал нетронутым — передавай копию.
>

## Указатели на срезы

```go
package main

import "fmt"

func main() {
	superpowers := &[3]string{"flight", "invisibility", "superStrength"}

	fmt.Printf("superpowers = %v\n", superpowers)

	fmt.Println(superpowers[0])

	fmt.Println(superpowers[1:])

	// Теперь для срезов
	//
	heroes := &[]string{"Me", "MyDog"}

	fmt.Printf("heroes = %v\n", heroes)
	fmt.Println((*heroes)[0]) // для срезов приходится писать так
	// fmt.Println(heroes[0]) // invalid operation: cannot index heroes (variable of type *[]string)

}
```

## Указатель на ресивер в методе

```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

// Метод с КОПИЕЙ (без указателя)
func (p person) birthdayCopy() {
	p.age++
}

// Метод с УКАЗАТЕЛЕМ (оригинал)
func (p *person) birthdayPointer() {
	p.age++
}

func main() {
	alice := person{name: "Alice", age: 20}

	// Изначально: {name:Alice age:20}
	fmt.Printf("Изначально: %+v\n", alice)

	// Метод, который работает с копией
	alice.birthdayCopy()
	// После birthdayCopy (копия): {name:Alice age:20}
	fmt.Printf("После birthdayCopy (копия): %+v\n", alice)

	// Метод, который работает с указателем
	alice.birthdayPointer()

	// После birthdayPointer (указатель): {name:Alice age:21}
	fmt.Printf("После birthdayPointer (указатель): %+v\n", alice)
}
```

| Метод | Что происходит | Оригинал меняется? |
| --- | --- | --- |
| `func (p person)` | Работает с копией | ❌ Нет |
| `func (p *person)` | Работает с оригиналом | ✅ Да |

## Разница в определении функций с указателем и без

Если метод реализован для значения — его можно вызывать как у обычного значения, так и у указателя. А вот если метод реализован для указателя, то использовать его можно только через указатель.

```go
package main

import (
	"fmt"
	"strings"
)

// Интерфейс talker требует один метод: talk() string
type talker interface {
	talk() string
}

// Функция shout принимает любой тип, реализующий интерфейс talker,
// вызывает его метод talk() и выводит результат в верхнем регистре.
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// ----- СЛУЧАЙ 1: Структура с методом по значению -----

// Тип messageValue — пустая структура (0 байт).
type messageValue struct{}

// Метод реализован для значения (без указателя):
// значит, интерфейс будет реализован и для messageValue{}, и для &messageValue{}.
func (m messageValue) talk() string {
	return "value interface speaking"
}

// ----- СЛУЧАЙ 2: Структура с методом по указателю -----

// Тип messagePointer — снова пустая структура.
type messagePointer struct{}

// Метод реализован только для *messagePointer (указателя).
// Значит, интерфейс реализуется только через &messagePointer{}.
func (m *messagePointer) talk() string {
	return "pointer interface speaking"
}

func main() {
	// ----- Пример с методом по значению -----
	fmt.Println("== Метод по значению ==")
	shout(messageValue{})   // ✅ Работает
	shout(&messageValue{})  // ✅ Тоже работает — Go сам разыменует

	// ----- Пример с методом по указателю -----
	fmt.Println("\n== Метод по указателю ==")
	// shout(messagePointer{})  // ❌ НЕ работает: нет метода talk у messagePointer
	shout(&messagePointer{})   // ✅ Работает: метод реализован у *messagePointer
}
```

Попытка обратиться к nil указателю вызывает панику.

Таким образом мы можем её избежать

```go
var nowhere *int
if nowhere != nil {
		fmt.Println(*nowhere)
}
```

# nil

Начальное значение для любого указателя это nil

Иногда пустой срез, map бывают взаимозаменяемы с nil

Например, если функция ожидает список или map, а нам нужно передать пустой объект, мы можем передать nil

```go
package main

import "fmt"

// Встроенные функции range, len и append не выдадут ошибку,
// если вместо пустого спика передать nil

func main() {
	var soup []string
	fmt.Println(soup == nil) // true
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup)) // 0
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup) // [onion carrot celery]
}

```

# Тесты

Как запустить тесты?

Запутить файл с тестами из основной директории

```go
go test
```

Запустить все тесты рекурсивно

```go
go test ./...
```

Как посмотреть покрытие кода?

Сначала генериурем отчёт

```go
go test -coverprofile=coverage.out
```

Потом можно создать html отчёт

```go
go tool cover -html=coverage.out -o coverage.html
```

В браузере ты увидишь исходный код с подсветкой:

- **Зелёные строки** — покрыты тестами.
- **Красные строки** — не покрыты.
- **Серые строки** — не исполняемые (комментарии, объявления).

# Переменная интерфейсного типа

В Go **интерфейс** — это **набор методов**, которые должен реализовать тип, чтобы "подходить под этот интерфейс". А переменная **интерфейсного типа** может хранить **любое значение**, главное — чтобы оно реализовывало нужные методы (а у interface{} — вообще никаких методов не требуется).

```go
var x interface{}
x = 42        // можно присвоить int
x = "hello"   // потом присвоить string
x = true      // и даже bool

```

Прикольно) За счёт interface{} можно делать переменные динамической типизации

## Что внутри переменной интерфейсного типа?

Когда ты присваиваешь значение интерфейсной переменной, Go **упаковывает два элемента**:

1. **Конкретное значение**
2. **Информацию о его типе (тип дескриптор)**

📦 Это называется **интерфейсная структура**: она не просто "содержит значение", а хранит и то, **что это за тип**, и **что это за значение**.

```go
var x interface{} = "Golang"

s, ok := x.(string)
if ok {
	fmt.Println("Это строка:", s)
} else {
	fmt.Println("Это не строка")
}

```

# Прикольные фишки

## Вывод аргумента и значения

```go
fmt.Printf("%#v", v")
```

Покажет сразу и значение и тип переменной

## Вывод полей пременной для стуркур и map

```go
fmt.Printf("%+v\n", *config)
```

## Обработка ошибок

Обёртка, чтобы проще обрабатывать ошибки

```go
type safeWriter struct {
	w   io.Writer // поток, куда пишем (например, файл)
	err error     // здесь храним первую ошибку, если она произойдёт
}

// Метод writeln записывает строку s, если раньше не было ошибки
func (sw *safeWriter) writeln(s string) {
	// Если уже была ошибка ранее, ничего не делаем
	if sw.err != nil {
		return
	}

	// Пишем строку. Если возникла ошибка, она сохраняется в sw.err
	_, sw.err = fmt.Fprintln(sw.w, s)
}
```

```go
package main

import (
	"bytes"
	"fmt"
	"io"
)

// safeWriter — это обёртка над io.Writer, которая сохраняет первую ошибку
type safeWriter struct {
	w   io.Writer // Куда писать
	err error     // Сюда запишется первая ошибка, если она возникнет
}

// writeln — метод записи строки. Если уже была ошибка, ничего не делает.
func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // Если ошибка уже была, выходим
	}
	_, sw.err = fmt.Fprintln(sw.w, s) // Пишем строку, запоминаем ошибку
}

func main() {
	var buf bytes.Buffer // Это "фейковый файл", просто буфер в памяти

	sw := safeWriter{w: &buf}

	sw.writeln("Первая строка")
	sw.writeln("Вторая строка")
	sw.writeln("Третья строка")

	// Проверим, была ли ошибка
	if sw.err != nil {
		fmt.Println("Ошибка при записи:", sw.err)
	} else {
		fmt.Println("Результат записи:")
		fmt.Println(buf.String()) // Печатаем, что получилось в буфере
	}
}
```

## type assertion

Можно использовать  **type assertion** в Go для извлечения конкретного типа ошибки из переменной типа error. Это полезно, когда ты возвращаешь пользовательский тип ошибки (например, срез ошибок) через интерфейс error, но хочешь получить доступ к его содержимому — например, распечатать каждую ошибку по отдельности.

```go
package main

import (
	"errors"
	"fmt"
)

// Объявляем пользовательский тип ошибки — срез из ошибок
type MyError []error

// Реализуем метод Error() для MyError, чтобы он соответствовал интерфейсу error
func (me MyError) Error() string {
	return fmt.Sprintf("%d errors occurred", len(me))
}

// Функция возвращает ошибку типа MyError, но как интерфейс error
func doSomething() error {
	return MyError{
		errors.New("first problem"),
		errors.New("second problem"),
	}
}

func main() {
	err := doSomething() // err имеет тип error, но внутри лежит MyError

	// Пытаемся привести err обратно к MyError с помощью type assertion
	if errs, ok := err.(MyError); ok {
		fmt.Printf("Произошло %d ошибок:\n", len(errs))
		for _, e := range errs {
			fmt.Println("-", e)
		}
	} else {
		// Если не удалось привести тип — просто выводим ошибку
		fmt.Println("Обычная ошибка:", err)
	}
}

// Произошло 2 ошибок:
// - first problem
// - second problem
```

## Цикл range

(начиная с go 1.22)

Вместо

```go
for i := 0; i < n; i++ {
    // тело цикла
}

```

Можно написать

```go
for i := range n {
    // тело цикла
}

```

Если переменная i не нужна, то

```go
for _ = range n {
    // повторить n раз без использования индекса
}

```

## Метки

```go
func main() {
    samples := []string{"hello", "apple_π!"}
outer:
    for _, sample := range samples {
        for i, r := range sample {
            fmt.Println(i, r, string(r))
            if r == 'l' {
                continue outer
            }
        }
    fmt.Println()
    }
}
```

## Опциональные параметры функции

Для имитации опциональных именованных параметров функции можно использовать структуры

```go
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	// выполнение каких-либо действий
}

func main() {
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}

```

## Переменное число параметров функции

```go
func variadicSum(message string, numbers ...int) int {
	fmt.Println(message)
	var sum int = 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(variadicSum("Hello world!", 1, 2, 3))
}

```

## Замыкания в defer

В Go оператор defer запоминает выражение функции, которое будет вызвано позже, но аргументы этого вызова вычисляются сразу, в момент выполнения строки defer.

```jsx
func Example() {
	i := 0
	defer fmt.Println(i) // 0
	i = 5
	fmt.Println("Here we go")
}
```

При выходе из функции  будет напечатано 0, потому что значение i было захвачено в момент создания отложенного вызова.

Но

Если же в отложенную функцию передаётся замыкание, а внутри него используется переменная из внешней области, то значение этой переменной будет определяться только при выполнении замыкания, т.е. после return

```jsx
func ExampleWithClosure() {
	i := 0
	defer func() {
		fmt.Println(i) // 5
	}()
	i = 5
	fmt.Println("Here we go")
}
```

Вывод будет 5, потому что внутри анонимного func() переменная i берётся из внешней области в момент выполнения отложенного кода, а не в момент создания defer.

```jsx
func main() {
	Example()     // 0
	fmt.Println() // 5
	ExampleWithClosure()
}

```

defer может читать и изменять возвращаемое значение, если имя возвращаемого значения задано в сигнатуре функции

```go
package main

import "fmt"

func compute() (result int) { // имя возвращаемого значения объявлено явно
	defer func() {
		if result == 0 {
			result = -1 // «ошибочный» код
		}
	}()

	val, ok := getValue()
	if !ok { // значение не получено – оставляем result = 0
		return
	}

	result = val
	return
}

// Имитируем получение данных
func getValue() (int, bool) {
	// return 42, true // <-- успешный вариант
	return 0, false // <-- вариант с ошибкой
}

func main() {
	fmt.Println("Успешный вызов:", compute())
}

```

## Указатель на указатель

```go
func update(g **int) {
    x := 10
    *g = &x
}

func main() {
    var f *int
    update(&f)
    fmt.Println(*f) // 10
}

```

Мы смогли внутри функции (неявно для основного кода) изменить значение с nil на *int

Если бы мы сделали вот так, была бы паника

```go
func failingUpdate2(g *int) {
	*g = 10 // panic
}

func main() {
	var f *int = nil

	failingUpdate2(f)
	fmt.Println(f)
}
```

В памяти

```go
main:
   f ───► nil

вызов failingUpdate2(f):
   g ───► nil (копия f)

   *g = 10
     │
     └── попытка записать по адресу nil ⇒ PANIC

```

Есть вот такой код

```go
func failingUpdate(g *int) {
    x := 10
    g = &x
}

func main() {
    var f *int = nil
    failingUpdate(f)
    fmt.Println(f) // nil
}

```

```go
main:
   f ───► nil

вызов failingUpdate(f):
   g ───► nil (копия f)
   x = 10 (лежит в стеке функции)

   g = &x   // теперь g указывает на x
   f остаётся nil, потому что g — копия!
```

## Кудрявые скобки без оператора

В Go можно ставить фигурные скобки без оператора

```go
// Ограничение видимости переменных
{
    x := 42
    fmt.Println(x)
}
// fmt.Println(x) // Ошибка: x не видно за пределами блока

```

```go
// defer который должен сработать до окончания функции
{
    f, _ := os.Open("file.txt")
    defer f.Close()

    // работа с файлом
}
// тут файл уже закрыт

```

```go
// Здесь это действительно полезно!

package main

import (
    "fmt"
    "os"
)

func main() {
    files := []string{"a.txt", "b.txt", "c.txt"}

    for _, name := range files {
        { // новый блок
            f, err := os.Create(name)
            if err != nil {
                panic(err)
            }
            defer f.Close() // сработает в конце БЛОКА, не всей функции
            fmt.Println("Created:", name)
        } // <-- здесь файл закроется
    }
    fmt.Println("Done")
}

```

```go
// Просто для группикровки логики и повышения
// читаемости
{
    fmt.Println("Start")
    fmt.Println("End")
}

```

## iota

Используется для создания последовательно возрастающих констант

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```

```go
const (
    A = iota + 1 // 1
    B             // 2
    C             // 3
)

```

```go
const (
    A = iota + 1 // 1
		_            // пропуск значения
    B             // 3
    C             // 4
)

```

Пример с битовыми флагами (когда нужно хранить несколько флагов (включено/выключено) в одном числе)

```go
package main

import "fmt"

const (
    Read = 1 << iota  // 1 << 0 = 1
    Write             // 1 << 1 = 2
    Execute           // 1 << 2 = 4
)

func main() {
    perms := Read | Execute       // комбинируем флаги
    fmt.Println(perms)            // 5
    fmt.Println(perms&Read != 0)  // true
    fmt.Println(perms&Write != 0) // false
}

// & проверяет, установлен ли конкретный бит
```

<< побитовый сдвиг влево

| Выражение | Результат в десятичной системе | В двоичной |
| --- | --- | --- |
| `1 << 0` | `1` | `0001` |
| `1 << 1` | `2` | `0010` |
| `1 << 2` | `4` | `0100` |

```go
// То есть
Read = 1
Write = 2
Execute = 4

```

Теперь у нас есть три флага, каждый в своём бите.
Их можно комбинировать побитовыми операциями:

```go
perms := Read | Execute
```

Пример 2

Дни недели

```go
package main

import "fmt"

type Weekday int // Создание типа, основанного на int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

// Создаёт массив строк, где индекс – значение iota
func (w Weekday) String() string {
    return [...]string{
        "Sunday", "Monday", "Tuesday", "Wednesday",
        "Thursday", "Friday", "Saturday",
    }[w]
}

func main() {
    fmt.Println(Monday)   // Monday
    fmt.Println(Friday)   // Friday
    fmt.Println(Weekday(6)) // Saturday
}

```

Пример 3

Направления

```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

// [...]string{"North", "East", "South", "West"}[d]
// Массив строк, берём элемент по индексу d
```

Поясниения

```go
func (w Weekday) String() string {
    return [...]string{
        "Sunday", "Monday", "Tuesday", "Wednesday",
        "Thursday", "Friday", "Saturday",
    }[w]
}

//"Если кто-то попытается вывести Weekday через fmt.Println() или %v,
// используй вот эту функцию, чтобы показать его в виде текста."
```

Уменьшающиеся последовательности

```go
const (
    A = 10 - iota // 10, 9, 8, 7
    B
    C
    D
)

```

```go
const (
    X = -iota // 0, -1, -2, -3
    Y
    Z
)

```

## Задаём свой формат вывода в Print и аналогах

Это можно сделать за счёт реализации метода String() (реализуем интерфейс)

```go
type Stringer interface {
    String() string
}
```

Цветной вывод

```go
package main

import "fmt"

type Status int

const (
	OK Status = iota
	Warning
	Error
)

func (s Status) String() string {
	switch s {
	case OK:
		return "\033[32mOK\033[0m" // зелёный
	case Warning:
		return "\033[33mWARNING\033[0m" // жёлтый
	case Error:
		return "\033[31mERROR\033[0m" // красный
	default:
		return "UNKNOWN"
	}
}

func main() {
	fmt.Println(OK, Warning, Error)
}

```

Пресонализованный вывод для структуры

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Admin bool
}

func (u User) String() string {
	data, _ := json.MarshalIndent(u, "", "  ")
	return string(data)
}

func main() {
	user := User{"Alice", 30, true}
	fmt.Println(user)
}

// Вывод
//{
//  "Name": "Alice",
//  "Age": 30,
//  "Admin": true
//}
// Вместо
// {Alice 30 true}

```

```go
package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point(x=%d, y=%d)", p.X, p.Y)
}

func main() {
	p := Point{10, 20}
	fmt.Println(p)
}

// Задали вывод
// Point(x=10, y=20)
```


Если реалиpовать метод GoString(), это можно будет использовать в форматированном выводе в %#v


```go

type Point struct{ X, Y int }
func (p Point) String() string   { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }
func (p Point) GoString() string { return fmt.Sprintf("Point{X:%d, Y:%d}", p.X, p.Y) }

p := Point{2,3}
fmt.Println(p)           // (2,3)
fmt.Printf("%#v\n", p)   // Point{X:2, Y:3}
```

## Alias типа

Можно создать не только тип на основе типа, а именно alias для основного типа

Даже если у типов одинаковый базовый тип, присваивание без преобразования запрещено:

```go

type MyInt int

var a int
var b MyInt

// a = b      // ошибка
a = int(b)    // ок: явное преобразование

type AliasInt = int
var c AliasInt
a = c         // ок: это тот же тип int (type alias)
```

## Измерение времени кода

```go
// Utility
func TrackTime(pre time.Time) time.Duration {
  elapsed := time.Since(pre)
  fmt.Println("elapsed:", elapsed)

  return elapsed
}

func TestTrackTime(t *testing.T) {
 defer TrackTime(time.Now()) // <--- THIS

 time.Sleep(500 * time.Millisecond)
}

// elapsed: 501.11125ms
```

## Итераторы 

Начиная с Go 1.23 появились итераторы

```go
package main

import (
	"fmt"
	"iter"
)

func countTo(s int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i <= s; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func countTo2(s int, divider int) iter.Seq2[int, bool] {
	// yield обязан принимать те же типы аргументов,
	// что и iter.Seq[A, B]
	return func(yield func(int, bool) bool) {
		for i := 0; i <= s; i++ {
			ok := i%divider != 0
			// Те же аргументы iter.Seq[A, B]
			if !yield(i, ok) {
				return
			}
		}
	}
}

func main() {
	for v := range countTo(5) {
		fmt.Println(v)
	}

	fmt.Println()

	for i, ok := range countTo2(7, 2) {
		fmt.Println(i, ok)
	}
}
```



# Без паники! (нет)

Как вызывать панику? (Звучит забавно)

```go
panic("Weeeee") // Сеем панику
```

## Что такое паника?

Паника — это **встроенный механизм ошибки**, который:

- прерывает обычное выполнение программы;
- начинает "раскрутку стека" — вызовы функций заканчиваются в обратном порядке;
- может быть **перехвачена** с помощью `recover()` (обычно в `defer`), если ты хочешь "спасти" программу от краха.

## Когда возникает паника?

### Go вызывает панику автоматически при:

- делении на ноль (`x / 0`);
- выходе за границы среза (`slice[100]`, если длина 10);
- разыменовании `nil`указателя;
- `runtime.Goexit()` (иногда);
- попытке разблокировать не заблокированный мьютекс (`sync.Mutex.Unlock()` без Lock);
- явном вызове `panic("что-то пошло не так")`.

Лучще обрабатывать ошибки через err ≠ nill, при этом painc лучше, чем os.Exit(), потому что при панике выполнятся deffer(), а при os.Exit() – нет

```go
package main

import (
	"fmt"
)

func testingSomething() {
	defer fmt.Println("GO")

	panic(0)
}

func main() {
	fmt.Println("Here we")

	testingSomething()
}

// Here we
// GO
// panic: 0
```

```go
// Без паники

package main

import (
	"fmt"
	"os"
)

func testingSomething() {
	defer fmt.Println("GO") // Не сработало

	os.Exit(0)
}

func main() {
	fmt.Println("Here we")

	testingSomething()
}

// Here we
```

## recover

recover позволяет перехватать панику и продолжить выполнение программы

```go
package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			// Этот код выполнится, если в main случится panic
			fmt.Println("Recover from panic:", e)
		}
	}()

	// Здесь вызывается паника
	panic("I forgot my towel")

	// Этот код никогда не будет выполнен
	fmt.Println("This line will not be printed")
}

```

```go
package main

import "fmt"

func riskyFunction() {
	fmt.Println("Начало riskyFunction")
	panic("Что-то пошло не так в riskyFunction!")
	// Эта строка уже не выполнится
	fmt.Println("Конец riskyFunction")
}

func main() {
	// Отложенная функция, которая перехватывает панику
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановление после паники:", r)
		}
	}()

	fmt.Println("Запуск riskyFunction")
	riskyFunction() // здесь произойдёт panic
	fmt.Println("После riskyFunction") // эта строка не выполнится
}

```

```go
package main

import "fmt"

func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("⚠️ Panic caught:", r)
			result = 0 // безопасное значение
		}
	}()

	return a / b // если b == 0 → panic
}

func main() {
	fmt.Println("10 / 2 =", safeDivide(10, 2))
	fmt.Println("10 / 0 =", safeDivide(10, 0)) // перехват panic
	fmt.Println("✅ Program continues normally")
}

```

# Gorutines

Если функция, вызвавшая gorutine завершается, то и горутины завершаются, даже если они не закончились

Такой код не выведет ничего

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i = i + 1 {
		go sleepyGopher()
	}

	// time.Sleep(10 * time.Second)
}

func sleepyGopher() {
	time.Sleep(4 * time.Second)
	fmt.Println("hrrrrr")
}

```

А если мы подождём

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range 5 {
		go sleepyGopher()
	}

	time.Sleep(10 * time.Second)
}

func sleepyGopher() {
	time.Sleep(4 * time.Second)
	fmt.Println("hrrrrr")
}

```

то уже выполняются все gorutines

При этом есть более элегантный способ подождать

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for _ = range 5 {
		wg.Add(1)
		go sleepyGopher(&wg)
	}

	wg.Wait() // ждём, пока все горутины закончат
}

func sleepyGopher(wg *sync.WaitGroup) {
	defer wg.Done() // сообщаем, что горутина завершена

	time.Sleep(4 * time.Second)
	fmt.Println("hrrrrr")
}

```

Бывают анонимными (объявляются с помощью слова Go)

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // небуферизованный канал

	go func() {
		ch <- 1 // блокируется, пока кто-то не прочитает
	}()

	x := <-ch // main горутина ждёт значение
	fmt.Println(x)
}

```

## wg.Wait(), wg.Go() и errgroup.Group()

Как делали раньше

```go
var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)  // Add BEFORE launching goroutine
        go func(id int) {
            defer wg.Done()  // Always use defer
            // Perform work
            fmt.Printf("Worker %d completed\\n", id)
        }(i)
    }
wg.Wait()  // Block until all complete
```

С обновлением 1.25

```go
var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Go(func() {
            // Work here - no manual Add()/Done()
            fmt.Printf("Worker %d completed\\n", i)
        })
    }
wg.Wait()
```

Реализация

```go
func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

```

## errgroup.Group

Это пакет `golang.org/x/sync/errgroup`.

Он делает то же самое, что `WaitGroup`, но **плюс**:

- умеет **собирать первую ошибку** из горутин;
- умеет работать с **контекстом** (`context.Context`).

```go
import "golang.org/x/sync/errgroup"

func main() {
	var g errgroup.Group

	for i := 0; i < 3; i++ {
		i := i
		g.Go(func() error {
			if i == 1 {
				return fmt.Errorf("worker %d failed", i)
			}
			fmt.Println("worker", i)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Ошибка:", err)
	}
}

```

# Каналы

### Небуфферизованный канал (по умолчанию)

```go
ch := make(chan int)
```

📌 **Что это значит**:

- Отправка `ch <- 42` **заблокирует горутину**, пока **другая горутина не прочтёт значение** `x := <-ch`.
- Аналогично: чтение из канала блокирует, если в канале пока ничего нет.

### Буфферизованный канал

```go
ch := make(chan int, 3) // буфер на 3 значения
```

📌 **Что это значит**:

- Можно отправить до 3 значений, **не блокируясь**.
- Чтение блокируется, если в буфере ничего нет.
- Отправка блокируется, если буфер **полный**.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // в буффере хранится до 2 значений
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch <- 3

	fmt.Println(<-ch)

	// fmt.Println(<-ch) // Ошибка будет, в канале пусто

}
```

### Однонаправленные каналы

Только на чтение или запись

```go
var sendOnly chan<- int     // только отправка
var receiveOnly <-chan int  // только чтение
```

Обычно используются в сигнатурах функций, чтобы ограничить поведение

```go
func producer(out chan<- int) {
    out <- 42
}

func consumer(in <-chan int) {
    fmt.Println(<-in)
}

func sendOnly(ch chan<- int) {
    ch <- 1
}

func readOnly(ch <-chan int) {
    fmt.Println(<-ch)
}
```

Ещё может использоваться

- В **API библиотеках** — чтобы указать, что функция только читает или пишет
- В **параллельных паттернах** — чтобы избежать ошибок, когда кто-то случайно пишет в канал, из которого должен только читать

🧠 **Итог**: однонаправленные каналы — это **ограничение доступа** на уровне компилятора, чтобы не допустить ошибок.

## Select

select позволяет одновременно ждать несколько операций с каналами и выполнить первую, которая станет доступна.

Это пример конкурентной программы, которая умеет ждать события и прерывать выполнение при превышении времени ожидания

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sleepyGopher — функция, которая "засыпает" на случайное время до 4 секунд,
// а затем отправляет свой id в канал
func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}

func main() {
	// Устанавливаем уникальный seed для генерации случайных чисел
	// Используем текущее время в наносекундах
	rand.Seed(time.Now().UnixNano())

	// Создаём небуферизованный канал для общения с горутинами
	c := make(chan int)

	// Запускаем 5 горутин, каждая будет "спать" и отправлять свой id
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	// Устанавливаем таймаут: сработает через 2 секунды
	timeout := time.After(2 * time.Second)

	// Ожидаем завершения 5 горутин или таймаута
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

```

Когда мы **перестаём ждать горутины** (например, сработал таймаут и мы вышли из `main`), это **не означает, что все горутины завершились**. Они продолжают выполняться **в фоне**, потребляя память и ресурсы.

Go **не завершает программу**, пока работает `main`-горутина. Но если `main()` закончилась, а другие горутины остались активными — **Go принудительно завершает процесс**, и такие горутины **обрываются посередине**, что может привести к утечкам ресурсов или другим багам.

### Твой компьютер в безопасности

Go-программа, даже если она завершилась с "висящими" горутинами, **не вредит компьютеру**. Она просто **завершается жёстко**, и:

- незавершённые горутины **обрываются**,
- **вся выделенная оперативная память освобождается**,
- никаких «утечек в систему» не происходит,
- **перезагрузка не требуется**.

### Что значит "может привести к утечкам ресурсов"?

Это **внутри самой программы**, не на уровне системы. Например:

- если ты открывал файл или соединение с базой данных, и не закрыл — это **внутренний баг** программы;
- если горутина зависла в бесконечном цикле или ждёт по каналу — **она просто не делает ничего полезного**, пока работает `main`.

### Что стоит делать как программисту:

- Использовать **каналы для завершения** (`done`), `context.WithCancel`, `WaitGroup` — чтобы **красиво завершать горутины**;
- Следить, чтобы **`main()` не завершалась, пока всё не завершено**;
- Если код учебный/демо — можно не заморачиваться. Но в настоящих приложениях — лучше практиковать правильную остановку.

## nil-каналы

Иногда ты используешь `select`, в котором участвуют **несколько каналов**, но **не все активны** одновременно. Тогда можно **"выключить" канал, присвоив ему nil**, и `select` будет просто игнорировать этот кейс.

Интересный пример с переключением каналов

```go
var ch chan int = nil

go func() {
	time.Sleep(2 * time.Second)
	ch = make(chan int, 1)
	ch <- 42
}()

for {
	select {
	case v := <-ch:
		fmt.Println("Got value:", v)
		return
	default:
		fmt.Println("Waiting...")
		time.Sleep(500 * time.Millisecond)
	}
}

```

🔍 Что происходит:

- `ch` сначала nil, поэтому `select` игнорирует `case v := <-ch`;
- Через 2 секунды `ch` становится рабочим;
- Тогда `select` "замечает" канал и получает значени

## Резюме

- **nil-канал** в Go — это канал, который никогда не сработает.
- Он **не вызывает панику**, но **блокирует отправку и получение навсегда**.
- Это поведение удобно использовать в `select`, чтобы **динамически "включать/выключать" каналы**.

Ещё один пример

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	dataCh := make(chan string)
	cancelCh := time.After(3 * time.Second) // канал срабатывает через 3 секунды

	go func() {
		// Имитация долгой работы
		time.Sleep(2 * time.Second)
		dataCh <- "Operation completed successfully"
	}()

	for {
		select {
		case msg := <-dataCh:
			fmt.Println("Success:", msg)
			// Операция завершилась — отключим cancelCh
			cancelCh = nil
			return
		case <-cancelCh:
			fmt.Println("Timeout! Operation took too long.")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

```

🧠 Что делает `cancelCh = nil`?

Это ключевой момент:

- Без этого `cancelCh` через 3 секунды **обязательно сработает**, даже если операция завершена.
- Если операция завершилась раньше, и мы **"обнулили" канал**, то `select` **перестаёт его слушать**.

🎓 В чём фишка nil-каналов в `select`?

- Если ты хочешь **гибко управлять логикой select**, ты можешь динамически включать/выключать нужные каналы с помощью `nil`.
- Это **чистый способ управления поведением без флагов, условий и повторяющегося кода.**

## Как завершать горутины красиво?

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sleepyGopher работает до тех пор, пока не придёт сигнал завершения
func sleepyGopher(id int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Gopher", id, "got the signal to stop.")
			return
		default:
			// имитация работы
			fmt.Println("Gopher", id, "is working...")
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	done := make(chan struct{}) // канал для остановки всех горутин

	for i := 0; i < 3; i++ {
		go sleepyGopher(i, done)
	}

	time.Sleep(2 * time.Second) // ждём немного
	fmt.Println("Main: time is up, asking goroutines to stop")

	close(done) // отправляем сигнал всем горутинам завершиться

	time.Sleep(1 * time.Second) // даём им время завершиться красиво
	fmt.Println("Main: all done")
}

```

# Generic

### Что такое generic в Go

**Generics** — это способ писать функции, типы или структуры, которые работают **с разными типами данных**, но при этом **без дублирования кода**.

📌 До Go 1.18 приходилось писать отдельную функцию под каждый тип (int, float64, string и т. д.). Теперь можно писать одну универсальную

```go
package main

import (
	"fmt"
)

// AddOrConcat обрабатывает только строго перечисленные типы: int, float64 и string
func AddOrConcat[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(AddOrConcat(3, 4))         // 7
	fmt.Println(AddOrConcat(1.5, 2.5))     // 4
	fmt.Println(AddOrConcat("Go", "lang")) // Golang

	// fmt.Println(AddOrConcat(true, false)) // ❌ ошибка компиляции — тип не поддерживается
}

```

Пример кода работы с каналами и горутинами

```go
package main

import (
	"fmt"
	"time"
)

// Эта функция спит 3 секунды, выводит сообщение и сообщает main-горутине,
// что она закончила, передавая свой id в канал.
func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id // Отправляем id обратно в канал
}

func main() {
	c := make(chan int) // создадим канал для коммуникации
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c) // запускаем функции параллельно
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // получаем из канала увемоление об окончании работы функции
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}

}

```

## Пример pipeline

Когда горутины передают значения от одной к другой

```go
package main

import (
	"fmt"
	"strings"
)

// Передаёт данные в другую горутину
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	downstream <- "" // Сигнал окончания передачи
}

// Фильтруен данные от прыдудущей и передаёт дальше
func filterGopher(upstream, downstream chan string) {
	for {
		item := <-upstream
		if item == "" {
			downstream <- "" // передаём сигнал дальше
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item // только хорошие строки
		}
	}
}

// Получает данные от предыдущей
// и печатает в консоль
func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)     // отправитель
	go filterGopher(c0, c1) // фильтрующий
	printGopher(c1)         // печатающий
}

```

- `sourceGopher` — передаёт строки по каналу.
- `filterGopher` — убирает строки, содержащие `"bad"`.
- `printGopher` — печатает строки на экран.
- Специальное значение `""` используется как **сигнал окончания работы**.
- Потоки соединены каналами: `c0` (передаёт от источника к фильтру), `c1` (от фильтра к принтеру).

При этом непонятно, передаётся ли пустая строка или мы таким образом сообщаем, что больше данных не будет. Лучше закрывать канал, когда больше не будем его использовать.

```go
close(c)
```

Если попробовать читать данные из закрытого канала, будем получать нулевые данные для любого типа. Если попробовать писать в закрытый канал, это вызовет панику

## Важный момент про закрытые каналы

Если читать по циклу из закрытого канала, получая нулевое значение, цикл будет бесконечным, что будет тратить много времени процессора.

Обязательно проверять, открыт ли канал перед чтением из него

Вот более безопасная версия этого кода

```go
package main

import (
	"fmt"
	"strings"
)

// Отправляет строки в канал и закрывает его
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream) // закрываем канал после отправки всех данных
}

// Фильтрует строки и передаёт только хорошие дальше, затем закрывает канал
func filterGopher(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			break // входной канал закрыт
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream) // закрываем выходной канал
}

// Получает и печатает данные из канала до его закрытия
func printGopher(upstream chan string) {
	for {
		v, ok := <-upstream
		if !ok {
			break // канал закрыт
		}
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

```

Альтернативный вариант

```go
package main

import (
	"fmt"
	"strings"
)

// Передаёт данные в канал и закрывает его
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream) // сигнал завершения
}

// Фильтрует данные и закрывает выходной канал
func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream) // закрываем после окончания чтения
}

// Получает и печатает значения, пока канал не закрыт
func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)     // генератор
	go filterGopher(c0, c1) // фильтрующий
	printGopher(c1)         // вывод
}

```

Здесь используется итерация по значениям канала, как альтернатива проверке, что канале не закрыт

# Для примера

## Номер один

```go
print("Hello Svetik!")
```

# Mutex

**Mutex (mutual exclusion — взаимное исключение)** — это механизм, который используется для **синхронизации доступа** к разделяемым (shared) данным между несколькими горутинами.

### Зачем он нужен?

Когда **несколько горутин одновременно** читают и/или записывают в **одну и ту же переменную**, может случиться **состояние гонки (race condition)**. Это приводит к неожиданным или неверным результатам.

**Мьютекс** позволяет сделать так, чтобы **только одна горутина одновременно** имела доступ к этим данным.

```go
package main

import "sync"

var mu sync.Mutex // создаём мьютекс

func main() {
	mu.Lock()         // захватываем мьютекс
	defer mu.Unlock() // освобождаем в конце функции

	// здесь может быть работа с общими данными
}

```

Как может выглядеть мьютекс

```go
type Visited struct {
	mu sync.Mutex            // защищает доступ к map
	visited map[string]int   // общая переменная
}
```

Это значит:

- В Go **методы и функции НЕ являются “потокобезопасными” по умолчанию**.
- Если ты **не видишь явной блокировки (mutex) или комментария “safe for concurrent use”**, **нельзя**использовать это из нескольких горутин одновременно.

- 🔒 Мьютекс — примитив для синхронизации доступа к общим данным.
- 👥 Если несколько горутин используют общие переменные — нужен мьютекс.
- 📦 Часто прячут мьютекс внутри структуры или пакета, чтобы защитить от неправильного использования.
- 💥 map в Go не потокобезопасна — при доступе из нескольких горутин нужна защита.

Примеры кода с мьютексами

```go
package main

import (
	"fmt"
	"sync"
)

// Visited безопасно отслеживает посещённые URL
type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

// NewVisited создает новую структуру с map
func NewVisited() *Visited {
	return &Visited{
		visited: make(map[string]int),
	}
}

// MarkVisited добавляет URL в map, защищённо
func (v *Visited) MarkVisited(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.visited[url]++
}

// HasVisited проверяет, был ли URL уже посещён
func (v *Visited) HasVisited(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	_, ok := v.visited[url]
	return ok
}

// Count возвращает общее количество посещений
func (v *Visited) Count() int {
	v.mu.Lock()
	defer v.mu.Unlock()

	return len(v.visited)
}

func main() {
	v := NewVisited()
	var wg sync.WaitGroup

	// Имитируем 10 параллельных посещений
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			url := fmt.Sprintf("https://example.com/page%d", i%3) // page0, page1, page2
			v.MarkVisited(url)
		}(i)
	}

	wg.Wait()

	fmt.Println("Общее количество уникальных URL:", v.Count())
	fmt.Println("Посещён ли page1?", v.HasVisited("https://example.com/page1"))
}

```

```go
package main

import (
	"fmt"
	"sync"
)

// Общая структура с числом и мьютексом
type Counter struct {
	mu    sync.Mutex
	value int
}

// Метод, безопасно увеличивающий счётчик
func (c *Counter) Increment() {
	c.mu.Lock()         // Захватываем мьютекс
	defer c.mu.Unlock() // Освобождаем при выходе из функции

	c.value++
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Запускаем 5 горутин
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение:", counter.value)
}

```

Интересный вопрос

> А где функция (горутина) проверят, открыт ли мьютекс? Можно ли начинать исполнение или мьютекс закрыт?
В каком месте кода?
>

Ответ

На самом деле **ты явно это не пишешь**. Это делает **внутри себя** функция `mu.Lock()`.

### Как работает mu.Lock()

- Когда ты вызываешь `mu.Lock()`, Go проверяет:
    - ❗ Есть ли уже кто-то, кто держит этот мьютекс?
        - Если **нет** — ты его получаешь сразу и идёшь дальше.
        - Если **да** — твоя горутина **блокируется** и **ждёт**, пока мьютекс освободят.
- После этого ты пишешь свой код, например, изменяешь общее значение.
- Когда ты вызываешь `mu.Unlock()`, мьютекс **освобождается**, и **одна** из ждущих горутин продолжает выполнение.

Однако есть метод если хочется попробовать сделать руками

```go
ok := mu.TryLock()
```

Он возвращает `true`, если захватить удалось, и `false`, если мьютекс занят. Но `TryLock()` появился только в **Go 1.18**, и используется реже.

## worker

## Что такое worker

Представь, что у тебя есть **рабочий**, который сидит и ждёт задания.

- Он **не делает ничего**, пока не поступит задача.
- Когда задача пришла — он **выполняет её**, затем снова **садится ждать**.
- Он не уходит домой, он работает бесконечно в цикле.

Этот рабочий — и есть **воркер**.

```go
func worker() {
    for {
        select {
        case job := <-jobChannel:
            // Обрабатываем задачу
        case <-quitChannel:
            // Кто-то сказал "хватит" — выходим
            return
        }
    }
}

```

Здесь хотят сказать: «**этот воркер будет вечно сидеть и ждать задачи из канала в select**»

```go
func worker() {
    for {
        select {
            // Wait for channels here.
        }
    }
}
```

🧠 Ещё проще: воркер = горутина с циклом, которая ждёт задачи и обрабатывает их.

# Мок

Вот код Gin сервера с подлючением к базе данных без интерфейсов

```go
package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func initDB() error {
	var err error
	// Подключение к PostgreSQL (измените строку подключения на свою)
	db, err = sql.Open("postgres", "host=localhost port=5432 user=your_user dbname=your_db sslmode=disable password=your_password")
	if err != nil {
		return err
	}

	return db.Ping()
}

func getDataHandler(c *gin.Context) {
	var result string

	// Простой запрос к базе данных
	err := db.QueryRow("SELECT 'Hello from database!' as message").Scan(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func main() {
	// Инициализация базы данных
	if err := initDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	// Создание роутера
	r := gin.Default()

	// Регистрация роута
	r.GET("/data", getDataHandler)

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
```

А вот с помощью интерфейсов. Можно мокать

```go
package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DatabaseInterface определяет интерфейс для работы с базой данных
type DatabaseInterface interface {
	GetMessage() (string, error)
}

// DatabaseService реализует DatabaseInterface
type DatabaseService struct {
	db *sql.DB
}

// Фабрика
func NewDatabaseService(db *sql.DB) *DatabaseService {
	return &DatabaseService{db: db}
}

func (ds *DatabaseService) GetMessage() (string, error) {
	var result string
	err := ds.db.QueryRow("SELECT 'Hello from database!' as message").Scan(&result)
	return result, err
}

// Handler содержит зависимости для обработчиков
type Handler struct {
	dbService DatabaseInterface
}

// Фабрика
func NewHandler(dbService DatabaseInterface) *Handler {
	return &Handler{dbService: dbService}
}

// Обработчик (handler)
func (h *Handler) getDataHandler(c *gin.Context) {
	result, err := h.dbService.GetMessage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func initDB() (*sql.DB, error) {
	// Подключение к PostgreSQL (измените строку подключения на свою)
	db, err := sql.Open("postgres", "host=localhost port=5432 user=your_user dbname=your_db sslmode=disable password=your_password")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Инициализация базы данных
	db, err := initDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	// Создание сервиса для работы с БД
	dbService := NewDatabaseService(db)

	// Создание обработчика с зависимостями
	handler := NewHandler(dbService)

	// Создание роутера
	r := gin.Default()

	// Регистрация роута
	r.GET("/data", handler.getDataHandler)

	// Запуск сервера на порту 8080
	r.Run(":8080")
}

// MockDatabaseService для тестирования
type MockDatabaseService struct {
	message string
	err     error
}

// Фабрика
func NewMockDatabaseService(message string, err error) *MockDatabaseService {
	return &MockDatabaseService{
		message: message,
		err:     err,
	}
}

func (mds *MockDatabaseService) GetMessage() (string, error) {
	return mds.message, mds.err
}

// Пример использования мока в тестах:
/*
func TestGetDataHandler(t *testing.T) {
	// Создаем мок
	mockDB := NewMockDatabaseService("Test message", nil)
	handler := NewHandler(mockDB)

	// Создаем gin контекст для тестирования
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Вызываем обработчик
	handler.getDataHandler(c)

	// Проверяем результат
	assert.Equal(t, http.StatusOK, w.Code)
	// ... другие проверки
}
*/
```

# Atomic

## Что делает пакет sync/atomic

Он обеспечивает **атомарность** — гарантирует, что операция выполнится **целиком**, без прерывания другими горутинами.

Это делается на уровне процессорных инструкций (очень быстро).

```go
import "sync/atomic"

var counter int64

func main() {
	// Безопасно прибавим значение
	atomic.AddInt64(&counter, 1) // counter += 1 (атомарно)

	// Безопасно читаем значение
	value := atomic.LoadInt64(&counter)

	// Безопасно записываем значение
	atomic.StoreInt64(&counter, 10)

	// Если значение равно 10, поменяй на 20
	ok := atomic.CompareAndSwapInt64(&counter, 10, 20)
	if ok {
    fmt.Println("успешно заменили 10 на 20")
	}
}
```

Пример использования

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Counter =", counter)
}

```

Без atomic результат был бы меньше 1000, были бы гонки

С atomic точно 1000

## Когда использовать

✅ Используй `sync/atomic`, если:

- нужно просто считать, инкрементировать или флаг включить/выключить;
- не хочется тянуть `sync.Mutex`.

❌ Не используй, если:

- у тебя сложная структура данных (тогда лучше `sync.Mutex` или `sync.RWMutex`);
- нужно обновлять **несколько переменных сразу** — `atomic` не поможет.

# Работа с зависимостями

Просмотр доступных обновлений зависимостей модуля

```go

go list -m -u all
```

Понять, зачем нужна зависимость

```go
go mod why -m mod
```

```go
go mod why -m example.com/dependency

example.com/dependency is imported by
example.com/yourmodule
```

# Init

Init функции выполняются до main. Причём сначала выполняются init функции в
импортируемых пакетах, потом в пакете main, потом уже функция main
