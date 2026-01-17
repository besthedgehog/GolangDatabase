package main

import (
	"testing"
)

// Тест 1: Проверка создания массива
func TestNewDynamicArray(t *testing.T) {
	// Кейс А: Нормальное создание
	da := NewDynamicArray[int](10)
	if da.Len() != 0 {
		t.Errorf("Expected len 0, got %d", da.Len())
	}
	if da.Cap() != 10 {
		t.Errorf("Expected cap 10, got %d", da.Cap())
	}

	// Кейс Б: Защита от отрицательной емкости
	daNegative := NewDynamicArray[int](-5)
	if daNegative.Cap() != 0 {
		t.Errorf("Expected cap 0 for negative input, got %d", daNegative.Cap())
	}
}

// Тест 2: Добавление элементов и автоматическое расширение (Resize)
func TestAddAndGrow(t *testing.T) {
	// Начинаем с емкости 2
	da := NewDynamicArray[int](2)

	// Добавляем 1-й и 2-й элементы (емкость полная)
	da.Add(10)
	da.Add(20)

	if da.Len() != 2 || da.Cap() != 2 {
		t.Errorf("Before grow: expected len=2, cap=2. Got len=%d, cap=%d", da.Len(), da.Cap())
	}

	// Добавляем 3-й элемент -> должно произойти удвоение (capacity 2 -> 4)
	da.Add(30)

	if da.Len() != 3 {
		t.Errorf("After add: expected len 3, got %d", da.Len())
	}
	if da.Cap() != 4 {
		t.Errorf("Expected capacity to double to 4, got %d", da.Cap())
	}

	// Проверяем, что данные не потерялись при копировании
	val, _ := da.Get(2)
	if val != 30 {
		t.Errorf("Expected value 30 at index 2, got %d", val)
	}
}

// Тест 3: Расширение с нулевой емкости (Edge case)
func TestGrowFromZero(t *testing.T) {
	da := NewDynamicArray[int](0)
	da.Add(100)

	if da.Len() != 1 {
		t.Errorf("Expected len 1, got %d", da.Len())
	}
	if da.Cap() < 1 {
		t.Errorf("Expected cap >= 1, got %d", da.Cap())
	}
	val, _ := da.Get(0)
	if val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}
}

// Тест 4: Получение элементов (Get) и ошибки границ
func TestGet(t *testing.T) {
	da := NewDynamicArray[string](2)
	da.Add("A")
	da.Add("B")

	// Успешный кейс
	val, err := da.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != "B" {
		t.Errorf("Expected 'B', got '%s'", val)
	}

	// Ошибочные кейсы (Table Driven Test)
	tests := []int{-1, 2, 100} // Индексы, которые должны вызвать ошибку
	for _, index := range tests {
		_, err := da.Get(index)
		if err == nil {
			t.Errorf("Expected error for index %d, but got nil", index)
		}
	}
}

// Тест 5: Обновление элементов (Put)
func TestPut(t *testing.T) {
	da := NewDynamicArray[int](2)
	da.Add(1)
	da.Add(2)

	// Меняем элемент по индексу 0
	err := da.Put(0, 999)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := da.Get(0)
	if val != 999 {
		t.Errorf("Expected 999, got %d", val)
	}

	// Проверка ошибки
	err = da.Put(5, 555)
	if err == nil {
		t.Error("Expected error when Putting out of bounds")
	}
}

// Тест 6: Удаление (Remove) и сдвиг
func TestRemove(t *testing.T) {
	da := NewDynamicArray[int](5)
	// Массив: [10, 20, 30, 40]
	da.Add(10)
	da.Add(20)
	da.Add(30)
	da.Add(40)

	// Удаляем индекс 1 (число 20). Ожидаем: [10, 30, 40]
	err := da.Remove(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Проверяем длину
	if da.Len() != 3 {
		t.Errorf("Expected len 3, got %d", da.Len())
	}

	// Проверяем сдвиг
	v0, _ := da.Get(0)
	v1, _ := da.Get(1) // Тут должно быть 30
	v2, _ := da.Get(2) // Тут должно быть 40

	if v0 != 10 || v1 != 30 || v2 != 40 {
		t.Errorf("Shift failed. Expected [10, 30, 40], got [%d, %d, %d]", v0, v1, v2)
	}

	// Проверяем ошибку при удалении неверного индекса
	err = da.Remove(10)
	if err == nil {
		t.Error("Expected error when removing out of bounds")
	}
}

// Тест 7: Проверка IsEmpty
func TestIsEmpty(t *testing.T) {
	da := NewDynamicArray[int](5)
	if !da.IsEmpty() {
		t.Error("New array should be empty")
	}

	da.Add(1)
	if da.IsEmpty() {
		t.Error("Array with elements should not be empty")
	}

	da.Remove(0)
	if !da.IsEmpty() {
		t.Error("Array should be empty after removing all elements")
	}
}

// Тест 8: Работа со структурами (Дженерики)
func TestWithStructs(t *testing.T) {
	da := NewDynamicArray[Dog](1)
	da.Add(Dog{"Rex", 5})
	da.Add(Dog{"Laika", 3})

	// Проверяем, что второй элемент записался корректно (с расширением памяти)
	dog, err := da.Get(1)
	if err != nil {
		t.Fatalf("Error getting dog: %v", err)
	}

	if dog.Name != "Laika" {
		t.Errorf("Expected Name Laika, got %s", dog.Name)
	}
}
