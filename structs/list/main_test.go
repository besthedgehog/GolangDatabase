package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPushTail(t *testing.T) {
// 	// Вставка перевого элемента
// 	t.Run("Push first element", func(t *testing.T) {
// 		list := NewSingleLinkedList[int]()
// 		list.PushTail(10)

// 		// Проверим длину
// 		if list.length != 1 {
// 			t.Errorf("expected len 1, found %v", list.length)
// 		}

// 		// Проверим, что список не пуст
// 		if list.IsEmpty() != false {
// 			t.Error("list should be empty")
// 		}

// 		// В списке из одного элемента head и tail это одно и то же
// 		if list.head != list.tail {
// 			t.Error("head and tail should be the same")
// 		}

// 		if list.head.data != 10 {
// 			t.Error("head should be 10")
// 		}

// 		if list.tail.data != 10 {
// 			t.Error("tail should be 10")
// 		}
// 	})

// 	// --- Сценарий 2: Вставка нескольких элементов (проверка связей) ---
// 	t.Run("Push multiple elements", func(t *testing.T) {
// 		list := NewSingleLinkedList[int]()

// 		// Добавляем 10, 20, 30
// 		list.PushTail(10)
// 		list.PushTail(20)
// 		list.PushTail(30)

// 		// 1. Проверяем итоговую длину
// 		if list.Size() != 3 {
// 			t.Errorf("Expected size 3, got %d", list.Size())
// 		}

// 		// 2. Проверяем Head (должен остаться 10)
// 		if list.head.data != 10 {
// 			t.Errorf("Expected head 10, got %d", list.head.data)
// 		}

// 		// 3. Проверяем Tail (должен стать 30)
// 		if list.tail.data != 30 {
// 			t.Errorf("Expected tail 30, got %d", list.tail.data)
// 		}

// 		// 4. Самое важное: ПРОВЕРЯЕМ ЦЕПОЧКУ (Traversing)
// 		// Мы вручную пройдемся по ссылкам nextPtr, чтобы убедиться, что вагоны сцеплены

// 		// 10 -> 20
// 		if list.head.nextPtr == nil {
// 			t.Fatal("Head.nextPtr is nil (link broken)")
// 		}
// 		if list.head.nextPtr.data != 20 {
// 			t.Errorf("Expected second node to be 20, got %d", list.head.nextPtr.data)
// 		}

// 		// 20 -> 30
// 		secondNode := list.head.nextPtr
// 		if secondNode.nextPtr == nil {
// 			t.Fatal("SecondNode.nextPtr is nil (link broken)")
// 		}
// 		if secondNode.nextPtr.data != 30 {
// 			t.Errorf("Expected third node to be 30, got %d", secondNode.nextPtr.data)
// 		}

// 		// 30 -> nil (Tail должен смотреть в никуда)
// 		thirdNode := secondNode.nextPtr
// 		if thirdNode.nextPtr != nil {
// 			t.Error("Tail.nextPtr should be nil")
// 		}
// 	})

// 	// --- Сценарий 3: Работа со строками (проверка Generics) ---
// 	t.Run("Works with strings", func(t *testing.T) {
// 		list := NewSingleLinkedList[string]()
// 		list.PushTail("Hello")
// 		list.PushTail("World")

// 		if list.head.data != "Hello" {
// 			t.Errorf("Expected Hello, got %s", list.head.data)
// 		}
// 		if list.tail.data != "World" {
// 			t.Errorf("Expected World, got %s", list.tail.data)
// 		}

// 		if list.head.nextPtr == nil {
// 			t.Fatal("Expected second node, but got nil")
// 		}

// 		if list.head.nextPtr.data != "World" {
// 			t.Errorf("Expected second node to be 'World', got '%s'", list.head.nextPtr.data)
// 		}
// 	})
// }

func TestSize(t *testing.T) {
	list := NewSingleLinkedList[int]()

	if list.Size() != 0 {
		t.Errorf("size should be 0, but size = %v", list.Size())
	}

	list.PushHead(1)
	if list.Size() != 1 {
		t.Errorf("size should be 1, but size = %v", list.Size())
	}
}

func TestEmpty(t *testing.T) {
	list := NewSingleLinkedList[string]()

	if list.IsEmpty() != true {
		t.Errorf("Expected list to be empty, but it's not")
	}

	list.PushTail("Hello")
	if list.IsEmpty() != false {
		t.Errorf("Expected list to be not empty, but it's empty")
	}
}

func TestPushTail(t *testing.T) {
	list := NewSingleLinkedList[int]()

	list.PushTail(10)

	if list.head.data != 10 {
		t.Errorf("Expected head to be 10, got %d", list.head.data)
	}
	if list.tail.data != 10 {
		t.Errorf("Expected tail to be 10, got %d", list.tail.data)
	}

	list.PushTail(17)

	if list.tail.data != 17 {
		t.Errorf("Expected tail to be 17, got %d", list.tail.data)
	}

	if list.head.data != 10 {
		t.Errorf("Expected head to be 10, got %d", list.head.data)
	}
}

func TestPushHead(t *testing.T) {
	list := NewSingleLinkedList[int]()

	list.PushHead(10)

	if list.head.data != 10 {
		t.Errorf("Expected head to be 10, got %d", list.head.data)
	}
	if list.tail.data != 10 {
		t.Errorf("Expected tail to be 10, got %d", list.tail.data)
	}

	list.PushHead(17)

	if list.head.data != 17 {
		t.Errorf("Expected head to be 17, got %d", list.head.data)
	}
	if list.tail.data != 10 {
		t.Errorf("Expected tail to be 10, got %d", list.tail.data)
	}
}

func TestGet(t *testing.T) {
	list := NewSingleLinkedList[int]()

	list.PushHead(1)
	list.PushHead(2)
	list.PushHead(3)

	{
		a, err := list.Get(0)
		assert.NoError(t, err, "should not be an error")
		assert.Equal(t, 3, a)
	}

	{
		b, err := list.Get(1)
		assert.NoError(t, err, "should not be an error")
		assert.Equal(t, 2, b)
	}

	{
		c, err := list.Get(2)
		assert.NoError(t, err, "should not be an error")
		assert.Equal(t, 1, c)
	}

}

// A -> B -> C
//
// A -> B -> X -> C
func TestInsert(t *testing.T) {
	list := NewSingleLinkedList[string]()

	list.PushTail("A") // 0
	list.PushTail("B") // 1
	list.PushTail("C") // 2

	list.Insert("X", 1)

	x, err := list.Get(1)
	assert.NoError(t, err, "should not be an error")
	assert.Equal(t, "X", x)

	{
		a, err := list.Get(0)
		assert.NoError(t, err, "should not be an error")
		assert.Equal(t, "A", a)
	}

	{
		x, err := list.Get(1)
		assert.NoError(t, err, "should not be an error")
		assert.Equal(t, "X", x)
	}

	{
		b, err := list.Get(2)
		assert.NoError(t, err, "should not be an error")
		assert.Equal(t, "B", b)
	}

	{
		c, err := list.Get(3)
		assert.NoError(t, err)
		assert.Equal(t, "C", c)
	}

}

func TestRemove(t *testing.T) {
	list := NewSingleLinkedList[int]()

	for i := 0; i <= 8; i++ {
		list.PushTail(i)
	}

	{
		err := list.Remove(99)
		assert.ErrorContains(t, err, "index out of range")
	}

	{
		err := list.Remove(-1)
		assert.ErrorContains(t, err, "index out of range")
	}

	{
		assert.Equal(t, 9, list.length)
		err := list.Remove(0)
		assert.NoError(t, err)
		assert.Equal(t, 8, list.length)
		assert.Equal(t, 1, list.head.data)
		assert.Equal(t, 8, list.tail.data)
	}

}

func TestForEach(t *testing.T) {
	list := NewSingleLinkedList[int]()

	var called bool = false
	list.ForEach(func(data int) {
		called = true
	})

	assert.False(t, called, "should not call fn on empty list")

	for i := 1; i <= 3; i++ {
		list.PushTail(i)
	}

	// собираем данные в список с помощью ForEach
	var collected []int
	list.ForEach(func(data int) {
		collected = append(collected, data)
	})

	assert.Equal(t, []int{1, 2, 3}, collected)
}
