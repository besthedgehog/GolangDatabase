package main

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
