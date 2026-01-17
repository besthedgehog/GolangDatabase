package main

import (
	"fmt"
)

type DynamicArray[T any] struct {
	arr []T
}

// Фабрика
func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	if capacity < 0 {
		capacity = 0
	}
	return &DynamicArray[T]{
		arr: make([]T, 0, capacity),
	}
}

// Getter
func (da *DynamicArray[T]) Len() int {
	return len(da.arr)
}

// Getter
func (da *DynamicArray[T]) Cap() int {
	return cap(da.arr)
}

// Проверка, является ли массив пустым
func (da *DynamicArray[T]) IsEmpty() bool {
	return da.Len() == 0
}

// Проверка на правильность обращения по индексу
// чтобы не было паники
func (da *DynamicArray[T]) checkRangeFromIndex(index int) error {
	if index >= da.Len() || index < 0 {
		return fmt.Errorf("Index %d out of range %d",
			index, da.Len())
	}
	return nil
}

// Выделение увеличение размера памяти, выделяемого под массив
func (da *DynamicArray[T]) newCapacity() {
	currentCap := da.Cap()
	var newCap int

	if currentCap == 0 {
		newCap = 1
	} else {
		newCap = currentCap * 2
	}

	newArr := make([]T, da.Len(), newCap)
	copy(newArr, da.arr)
	da.arr = newArr
}

// Удаление элемента. Смещаем все элементы на один индекс влево
func (da *DynamicArray[T]) Remove(index int) error {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return err
	}

	copy(da.arr[index:], da.arr[index+1:])
	// Пояснение
	// dst:   [2 3 4]
	// src:   [3 4]
	// res: [1 3 4 4]

	// Теперь узбавимся от последнего элемента (дубликата)
	// Важный момент для Garbage Collector'а в Go!
	// Последний элемент теперь дублируется, его нужно "занулить",
	// иначе ссылка на объект останется в памяти, даже если мы уменьшим len.
	var zero T
	da.arr[len(da.arr)-1] = zero
	da.arr = da.arr[:len(da.arr)-1]
	return nil
}

// Получить элемент по индексу
func (da *DynamicArray[T]) Get(index int) (T, error) {
	if err := da.checkRangeFromIndex(index); err != nil {
		var zero T
		return zero, err
	}
	return da.arr[index], nil
}

// Обновление элемента массива
func (da *DynamicArray[T]) Put(index int, element T) error {
	if err := da.checkRangeFromIndex(index); err != nil {
		return err
	}
	da.arr[index] = element
	return nil
}

// Экземпляры этой структуры будут элементами массива
type Dog struct {
	Name string
	Age  uint8
}

func (da *DynamicArray[T]) Add(element T) {

	if da.Len() == da.Cap() {
		da.newCapacity()
	}

	// Увеличим длину массива на 1, чтобы не было паники при добавлении
	da.arr = da.arr[:len(da.arr)+1]
	// добавим элемент
	da.arr[len(da.arr)-1] = element
}

func main() {
	// Создаем массив с емкостью 1
	dogs := NewDynamicArray[Dog](1)

	fmt.Printf("Start: Len=%d, Cap=%d\n", dogs.Len(), dogs.Cap())

	// Добавляем собак
	dogs.Add(Dog{"Rex", 5})
	fmt.Printf("Added Rex: Len=%d, Cap=%d\n", dogs.Len(), dogs.Cap())

	dogs.Add(Dog{"Bobik", 3}) // Тут произойдет grow (1 -> 2)
	fmt.Printf("Added Bobik: Len=%d, Cap=%d\n", dogs.Len(), dogs.Cap())

	dogs.Add(Dog{"Laika", 2}) // Тут произойдет grow (2 -> 4)
	fmt.Printf("Added Laika: Len=%d, Cap=%d\n", dogs.Len(), dogs.Cap())

	fmt.Println("--- Список собак ---")
	for i := 0; i < dogs.Len(); i++ {
		dog, _ := dogs.Get(i)
		fmt.Printf("%d: %s (%d years)\n", i, dog.Name, dog.Age)
	}

	// Удаляем Бобика (индекс 1)
	fmt.Println("--- Удаляем индекс 1 ---")
	err := dogs.Remove(1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := 0; i < dogs.Len(); i++ {
		dog, _ := dogs.Get(i)
		fmt.Printf("%d: %s\n", i, dog.Name)
	}

	fmt.Printf("End state: Len=%d, Cap=%d\n", dogs.Len(), dogs.Cap())
}
