package main

// Node узел в кольцевом списке
type Node[T any] struct {
	data    T
	nextPtr *Node[T]
	prevPtr *Node[T]
}

// Dog структура для примера
type Dog struct {
	name string
	age  uint8
}

// CyclicList кольцевой список
type CyclicLinkedList[T any] struct {
	length int
	head   *Node[T]
}

// newNode возвращает указатель на ноду
func newNode[T any](data T) *Node[T] {
	return &Node[T]{
		data:    data,
		nextPtr: nil,
		prevPtr: nil,
	}
}

func NewCyclicLinkedList[T any]() *CyclicLinkedList[T] {
	return &CyclicLinkedList[T]{length: 0, head: nil}
}

func (cl *CyclicLinkedList[T]) Size() int {
	return cl.length
}

func (cl *CyclicLinkedList[T]) IsEmpty() bool {
	return cl.length == 0
}

// НАПИСАТЬ ТЕСТ
func (cl *CyclicLinkedList[T]) Add(data T) {

	// newNode уже возвращает указатель на ноду
	B := newNode(data)

	if cl.IsEmpty() {
		// Новая нода теперь вершина
		cl.head = B
		{
			B.nextPtr = B
			B.prevPtr = B

		}

		cl.length++
		return
	}

	// Мы встраиваем B между A и C
	//    B
	//    |
	// A <-> C
	//

	// head
	// ↓
	// A <-> ... <-> C
	// ↑             ↓
	// ←←←←←←←←←←←←←←←

	//
	//		head
	//      ↓
	// A <-> B    <-> C <->
	// ↑                    ↓
	// ←←←←←←←←←←←←←←←←←←←←←

	A := cl.head   // Текущая голова
	C := A.prevPtr // Текущий хвост

	// Новая нода встанет за А

	// Сначала с B
	B.nextPtr = A
	B.prevPtr = C
	// B.prevPtr = A.prevPtr

	// C.prevPtr  // не меняется
	C.nextPtr = B

	// A.nextPtr // не меняется
	A.prevPtr = B

	cl.head = B
	cl.length++
	return
}

// ForEach применяет функцию fn к каждом элементу кольцевого списка
func (cl *CyclicLinkedList[T]) ForEach(fn func(data T)) {
	node := cl.head

	fn(node.data)
	for range cl.length - 1 {
		node = node.nextPtr
		fn(node.data)
	}
}

func (cl *CyclicLinkedList[T]) ReverseForEach(fn func(data T)) {
	node := cl.head

	fn(node.data)
	for range cl.length - 1 {
		node = node.prevPtr
		fn(node.data)
	}
}
