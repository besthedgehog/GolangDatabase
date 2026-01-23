package main

import "errors"

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
	if index < 0 || index >= dl.length {
		return errors.New("index out the range")
	}

	if index == 0 {
		return dl.PushHead(data)
	}

	if index == dl.length-1 {
		return dl.PushTail(data)
	}

	node := dl.head
	for it := 0; it < index; it++ {
		node = node.nextPtr
	}

	insertNode := NewNode(data)
	insertNode.nextPtr = node
	node.prevPtr.nextPtr = insertNode
	insertNode.prevPtr = node.prevPtr
	node.prevPtr = insertNode

	dl.length++
	return nil
}
