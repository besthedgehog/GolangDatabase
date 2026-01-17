package main

import (
	"fmt"
)

type DynamicArray[T any] struct {
	length   int
	capacity int
	arr      []T
}

// Фабрика
func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	return &DynamicArray[T]{
		arr: make([]T, 0, capacity),
	}
}

// Проверка на правильность обращения по индексу
// чтобы не было паники
func (da *DynamicArray[T]) checkRangeFromIndex(index int) error {
	if index >= da.length || index < 0 {
		return fmt.Errorf("Index %d out of range %d",
			index, da.length)
	}
	return nil
}

// Выделение увеличение размера памяти, выделяемого под массив
func (da *DynamicArray[T]) newCapacity() {
	da.capacity *= 2
	newArr := make([]T, da.length, da.capacity)
	copy(newArr, da.arr)
	da.arr = newArr
}

// Проверка, является ли массив пустым
func (da *DynamicArray[T]) IsEmpty() bool {
	return da.length == 0
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
	var zero T
	da.arr[len(da.arr)-1] = zero
	da.arr = da.arr[:len(da.arr)-1]
	return nil
}

// Получить элемент по индексу
func (da *DynamicArray[T]) Get(index int) (T, error) {
	if err := da.checkRangeFromIndex(index); err != nil {
		return *new(T), err
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
	name string
	age  uint8
}

func (da *DynamicArray[T]) Add(element T) {
	if len(da.arr) == cap(da.arr) {
		newArr := make([]T, len(da.arr), cap(da.arr)*2)
		copy(newArr, da.arr)
		da.arr = newArr
	}

	// Увеличим длину массива на 1, чтобы не было паники при добавлении
	da.arr = da.arr[:len(da.arr)+1]
	// добавим элемент
	da.arr[len(da.arr)-1] = element
}

func main() {
	dynamicArray := NewDynamicArray[Dog](1)
	fmt.Println(dynamicArray)
	dynamicArray.Add(Dog{"Max", 4})
	// dynamicArray.Add(Dog{"Alex", 3})
	// dynamicArray.Add(Dog{"Misha", 1})
	fmt.Println(dynamicArray)

}
