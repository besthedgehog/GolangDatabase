package main

import (
	"fmt"
)

type DynamicArray[T any] struct {
	length   int
	capacity int
	arr      []T
}

// func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
// 	if capacity <= 0 {
// 		panic("Array capacity cannot be <= 0")
// 	}

// 	return &DynamicArray[T]{
// 		capacity: capacity,
// 		arr:      make([]T, capacity),
// 	}
// }

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
	da.capacity = da.capacity * 2

	newArr := make([]T, da.capacity)
	copy(newArr, da.arr)
	da.arr = newArr
}

// Проверка, является ли массив пустым
func (da *DynamicArray[T]) IsEmpty() bool {
	return da.length == 0
}

// // Добавление нового элемента в массив
// func (da *DynamicArray[T]) Add(element T) {
// 	if da.length == da.capacity {
// 		da.newCapacity()
// 	}

// 	da.arr[da.length] = element
// 	da.length++
// }

// Удаление элемента. Смещаем все элементы на один индекс влево
func (da *DynamicArray[T]) Remove(index int) error {
	err := da.checkRangeFromIndex(index)
	if err != nil {
		return err
	}

	copy(da.arr[index:], da.arr[index+1:])

	// Делаем последний элемент нулевым, чтобы не было утечек памяти
	da.arr[da.length-1] = *new(T) // зачем
	da.length--
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

func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	return &DynamicArray[T]{
		arr: make([]T, 0, capacity),
	}
}

func (da *DynamicArray[T]) Add(element T) {
	if len(da.arr) == cap(da.arr) {
		newArr := make([]T, len(da.arr), cap(da.arr)*2)
		copy(newArr, da.arr)
		da.arr = newArr
	}

	da.arr = da.arr[:len(da.arr)+1]
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
