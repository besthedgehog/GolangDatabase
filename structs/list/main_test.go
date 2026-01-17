package main

import "testing"

func TestPushTail(t *testing.T) {
	// Вставка перевого элемента
	t.Run("Push first element", func(t *testing.T) {
		list := NewSingleLinkedList[int]()
		list.PushTail(10)

		// Проверим длину
		if list.length != 1 {
			t.Errorf("expected len 1, found %v", list.length)
		}

		// Проверим, что список не пуст
		if list.IsEmpty() != false {
			t.Error("list should be empty")
		}

		// В списке из одного элемента head и tail это одно и то же
		if list.head != list.tail {
			t.Error("head and tail should be the same")
		}

		if list.head.data != 10 {
			t.Error("head should be 10")
		}

		if list.tail.data != 10 {
			t.Error("tail should be 10")
		}
	})

	// --- Сценарий 2: Вставка нескольких элементов (проверка связей) ---
	t.Run("Push multiple elements", func(t *testing.T) {
		list := NewSingleLinkedList[int]()

		// Добавляем 10, 20, 30
		list.PushTail(10)
		list.PushTail(20)
		list.PushTail(30)

		// 1. Проверяем итоговую длину
		if list.Size() != 3 {
			t.Errorf("Expected size 3, got %d", list.Size())
		}

		// 2. Проверяем Head (должен остаться 10)
		if list.head.data != 10 {
			t.Errorf("Expected head 10, got %d", list.head.data)
		}

		// 3. Проверяем Tail (должен стать 30)
		if list.tail.data != 30 {
			t.Errorf("Expected tail 30, got %d", list.tail.data)
		}

		// 4. Самое важное: ПРОВЕРЯЕМ ЦЕПОЧКУ (Traversing)
		// Мы вручную пройдемся по ссылкам nextPtr, чтобы убедиться, что вагоны сцеплены

		// 10 -> 20
		if list.head.nextPtr == nil {
			t.Fatal("Head.nextPtr is nil (link broken)")
		}
		if list.head.nextPtr.data != 20 {
			t.Errorf("Expected second node to be 20, got %d", list.head.nextPtr.data)
		}

		// 20 -> 30
		secondNode := list.head.nextPtr
		if secondNode.nextPtr == nil {
			t.Fatal("SecondNode.nextPtr is nil (link broken)")
		}
		if secondNode.nextPtr.data != 30 {
			t.Errorf("Expected third node to be 30, got %d", secondNode.nextPtr.data)
		}

		// 30 -> nil (Tail должен смотреть в никуда)
		thirdNode := secondNode.nextPtr
		if thirdNode.nextPtr != nil {
			t.Error("Tail.nextPtr should be nil")
		}
	})

	// --- Сценарий 3: Работа со строками (проверка Generics) ---
	t.Run("Works with strings", func(t *testing.T) {
		list := NewSingleLinkedList[string]()
		list.PushTail("Hello")
		list.PushTail("World")

		if list.head.data != "Hello" {
			t.Errorf("Expected Hello, got %s", list.head.data)
		}
		if list.tail.data != "World" {
			t.Errorf("Expected World, got %s", list.tail.data)
		}

		if list.head.nextPtr == nil {
			t.Fatal("Expected second node, but got nil")
		}

		if list.head.nextPtr.data != "World" {
			t.Errorf("Expected second node to be 'World', got '%s'", list.head.nextPtr.data)
		}
	})
}
