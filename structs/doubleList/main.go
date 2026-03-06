package main

import (
	"errors"
	"fmt"
)

// Node[T] узел в двусвязном списке
type Node[T any] struct {
	data    T
	nextPtr *Node[T]
	prevPtr *Node[T]
}

// DoublyLinkedList[T] структура, которая представляет сам список
type DoublyLinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

// NewNode фабрика узла
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{data: data, nextPtr: nil, prevPtr: nil}
}

// NewDoublyLinkedList фабрика двусвязного списка
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{length: 0, head: nil, tail: nil}
}

// Size возвращает длину списка
func (dl *DoublyLinkedList[T]) Size() int {
	return dl.length
}

// PushTail добавляет в конец списка
func (dl *DoublyLinkedList[T]) PushTail(data T) error {
	node := NewNode[T](data)

	if dl.length == 0 {
		dl.head = node
		dl.tail = node
		dl.length = 1
		return nil
	}

	dl.tail.nextPtr = node
	node.prevPtr = dl.tail
	dl.tail = node
	dl.length++
	return nil
}

func (dl *DoublyLinkedList[T]) PushHead(data T) error {

	node := NewNode(data)

	if dl.length == 0 {
		dl.head = node
		dl.tail = node
		dl.length++
		return nil
	}

	node.nextPtr = dl.head
	dl.head.prevPtr = node
	dl.head = node
	dl.length++
	return nil
}

// Insert добавляет элемент по индексу
func (dl *DoublyLinkedList[T]) Insert(index int, data T) error {
	if index < 0 || index > dl.length {
		return errors.New("index out the range")
	}

	if index == 0 {
		return dl.PushHead(data)
	}

	if index == dl.length {
		return dl.PushTail(data)
	}

	node := dl.head
	// change using range over int
	// for it := 0; it < index; it++ {
	// 	node = node.nextPtr
	// }

	for range index {
		node = node.nextPtr
	}

	// 0                    1               2
	// [node.prevPrt] <-> [node] <-> [node.nextPrt]
	// Insert(1)
	//
	// 			   [insertNode]
	//                  |
	// 					|
	// [node.prevPrt] <-> [node] <-> [node.nextPrt]
	//

	insertNode := NewNode(data)
	insertNode.nextPtr = node

	node.prevPtr.nextPtr = insertNode

	insertNode.prevPtr = node.prevPtr
	node.prevPtr = insertNode

	dl.length++
	return nil
}

// Get возвращает по индексу элемент
func (dl *DoublyLinkedList[T]) Get(index int) (T, error) {
	if index < 0 || index >= dl.length {
		var zeroT T
		return zeroT, errors.New("index out of range")
	}

	if index == 0 {
		return dl.head.data, nil
	}

	if index == dl.length-1 {
		return dl.tail.data, nil
	}

	node := dl.head

	for range index {
		node = node.nextPtr
	}

	return node.data, nil
}

// // Remove удаляет элемент по индексу
// func (dl *DoublyLinkedList[T]) Remove(index int) error {
// 	if index < 0 || index >= dl.length {
// 		return errors.New("index out of range")
// 	}

// 	{
// 		if index == 0 {
// 			node := dl.head
// 			dl.head = node.nextPtr
// 			dl.head.prevPtr = nil
// 			dl.length--
// 			return nil
// 		}
// 	}

// 	node := dl.head
// 	for it := 0; it < index-1; it++ {
// 		node = node.nextPtr
// 	}

// 	{
// 		if index == dl.length-1 {
// 			dl.tail.prevPtr = nil
// 			dl.tail = node
// 			dl.tail.nextPtr = nil
// 			dl.length--
// 			return nil
// 		}
// 	}

// 	// убрать этот кусок
// 	{
// 		B := node
// 		_ = B

// 		// deleteNode := B.nextPrt
// 		deleteNode := node.nextPtr
// 		C := deleteNode
// 		_ = C

// 		node.nextPtr = deleteNode.nextPtr // node.nextPtr = node.nextPtr.nextPtr (эквивалентно)
// 		// B.nextPtr = C.nextPtr // сразу указали с B на D

// 		node.nextPtr.prevPtr = deleteNode.prevPtr
// 		// D.prevPtr = C.prevPrt

// 		// «у ноды, следующей за удаляемой, указатель prevPtr
// 		// теперь должен указывать туда же, куда указывал prevPtr удаляемой ноды»
// 	}

// 	// Нода для удаления
// 	deleteNode := node.nextPtr

// 	// 0                    1               2
// 	// [node.prevPrt] <-> [node] <-> [node.nextPrt] <-> [node.nextPtr.nextPtr]
// 	//
// 	//
// 	// 			  					 [deletetNode]
// 	//                  					|
// 	// 										|
// 	// [node.prevPrt] <-> [node] <-> [node.nextPrt] <-> [node.nextPtr.nextPtr]
// 	//
// 	// [node.prevPtr] <-> [node] <-> [deleteNode] <-> [deleteNode.nextPtr]

// 	node.nextPtr = deleteNode.nextPtr // node.nextPtr = node.nextPtr.nextPtr

// 	//
// 	//
// 	// «у ноды, следующей за удаляемой, указатель prevPtr
// 	// теперь должен указывать туда же, куда указывал prevPtr удаляемой ноды»

// 	node.nextPtr.prevPtr = deleteNode.prevPtr // node.nextPtr.prevPrt = node.nextPtr.prevPtr (?????)

// 	//
// 	dl.length--
// 	return nil
// }

// A    <-> B    <-> C   <-> D
// prev <-> node <-> del <-> next
func (dl *DoublyLinkedList[T]) Remove(index int) error {
	if index < 0 || index >= dl.length {
		return errors.New("index out of range")
	}

	if index == 0 {
		A := dl.head // его удаляем
		B := A.nextPtr
		B.prevPtr = nil // отсоединили A =! B
		dl.head = B
		dl.length--
		return nil
	}

	if index == dl.length-1 {
		D := dl.tail
		C := D.prevPtr
		C.nextPtr = nil // С !!= D
		dl.tail = C
		dl.length--
		return nil
	}

	node := dl.head // A
	for i := 0; i < index-1; i++ {
		node = node.nextPtr
	}

	// Удалем 2 элемент
	// node = A,
	//
	// i = 0 < (2-1) {
	// 	node = B
	//	}
	//
	// i = 1 = (2-1) {цикл законен}

	// допустим i = 2
	B := node

	C := B.nextPtr

	D := C.nextPtr

	B.nextPtr = D

	fmt.Println(D)

	D.prevPtr = B

	dl.length--

	return nil
}

func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	if dl.length == 0 {
		return true
	}
	return false
}

// ForEach применяет функцию для каждого элемента
func (dl *DoublyLinkedList[T]) ForEach(fn func(data T)) {
	node := dl.head
	for node != nil {
		fn(node.data)
		node = node.nextPtr
	}
}

func (dl *DoublyLinkedList[T]) ForEachReverse(fn func(data T)) {
	node := dl.tail
	for node != nil {
		fn(node.data)
		node = node.prevPtr
	}
}
