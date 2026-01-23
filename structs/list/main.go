package main

import (
	"errors"
	"fmt"
)

// 		 HEAD                                TAIL
//         ↓                                   ↓
//       [10]  ---->  [20]  ---->  [30]  ----> nil
//       (data)      (next)

// Односвязный список
type Node[T any] struct {
	data    T
	nextPtr *Node[T]
}

// Фабрика
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		data:    data,
		nextPtr: nil,
	}
}

// Весь список целиком
type SingleLinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

// Фабрика
func NewSingleLinkedList[T any]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (sl *SingleLinkedList[T]) Size() int {
	return sl.length
}

func (sl *SingleLinkedList[T]) IsEmpty() bool {
	return sl.length == 0
}

// Добавление в конец списка
func (sl *SingleLinkedList[T]) PushTail(data T) error {
	// Создаём узел
	node := NewNode[T](data)

	// Если список пуст
	if sl.length == 0 {
		sl.head = node
		sl.tail = node
		sl.length = 1
		return nil
	}

	// Это поле внутри старого последнего узла
	sl.tail.nextPtr = node

	// Это поле внутри нашей структуры SingleLinkedList
	sl.tail = node

	sl.length++
	return nil
}

// Добавление с головы
func (sl *SingleLinkedList[T]) PushHead(data T) error {
	node := NewNode[T](data)

	// Если список пуст
	if sl.length == 0 {
		sl.head = node
		sl.tail = node
		sl.length = 1
		return nil
	}

	node.nextPtr = sl.head
	sl.head = node
	sl.length++
	return nil
}

func (sl *SingleLinkedList[T]) Insert(data T, index int) error {
	if index < 0 || index >= sl.length {
		return errors.New("index out of range")
	}

	if index == 0 {
		return sl.PushHead(data)
	}

	if index == sl.length-1 {
		return sl.PushTail(data)
	}

	// создали ноду
	node := sl.head
	// доехали до ноды ДО той, где вставка будет
	for it := 0; it < index-1; it++ {
		node = node.nextPtr
	}

	// новая нода
	insertNode := NewNode[T](data)
	// Прицепляем новую ноду к следующей
	insertNode.nextPtr = node.nextPtr
	// а текущую ноду к новой
	node.nextPtr = insertNode

	sl.length++
	return nil
}

// Get возвращает объект по индексу и ошибку
func (sl *SingleLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= sl.length {
		var zero T
		return zero, errors.New("index out of range")
	}

	if index == 0 {
		return sl.head.data, nil
	}

	if index == sl.length-1 {
		return sl.tail.data, nil
	}

	node := sl.head
	for it := 0; it < index; it++ {
		node = node.nextPtr
	}

	return node.data, nil
}

// Remove удаляет элемент по индексу
func (sl *SingleLinkedList[T]) Remove(index int) error {
	if index < 0 || index >= sl.length {
		return errors.New("index out of range")
	}

	// Тут точно норм?
	if index == 0 {
		sl.head = sl.head.nextPtr
		if sl.head == nil {
			sl.tail = nil
		}
		sl.length--
		return nil
	}

	node := sl.head
	for it := 0; it < index-1; it++ {
		node = node.nextPtr
	}

	if index == sl.length-1 {
		sl.tail = node // выкинули последний элемент
		node.nextPtr = nil
		sl.length--
		return nil
	}

	// deleteNode := node.nextPtr
	// node.nextPtr = deleteNode.nextPtr
	node.nextPtr = node.nextPtr.nextPtr
	sl.length--
	return nil
}

// ForEach принимает на вход функцию, вызываемую для каждого элемента списка
func (sl *SingleLinkedList[T]) ForEach(fn func(data T)) {
	node := sl.head
	if node == nil {
		return
	}
	fn(node.data)
	for node.nextPtr != nil {
		node = node.nextPtr
		fn(node.data)
	}
}

// На при мере это структуры сделаем список
type Dog struct {
	name string
	age  int
}

func main() {
	sLinkedList := NewSingleLinkedList[Dog]()
	sLinkedList.PushTail(Dog{name: "Buddy", age: 3})
	sLinkedList.PushTail(Dog{name: "Max", age: 5})
	sLinkedList.PushTail(Dog{name: "Charlie", age: 2})

	sLinkedList.ForEach(func(dog Dog) {
		fmt.Printf("Name: %s, Age: %d\n", dog.name, dog.age)
	})
}
