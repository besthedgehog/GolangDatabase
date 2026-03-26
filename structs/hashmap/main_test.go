package main

import (
	"testing"
)

// func TestSomething(t *testing.T) {
// 	var a uint64 = 10000000

// 	fmt.Println(a)
// 	fmt.Println()
// 	b := a >> 3
// 	fmt.Println(b)
// 	c := a << 1
// 	fmt.Println(c)
// }

// func TestCheckMath(t *testing.T) {
// 	var hash uint64 = 123456789
// 	var capacity uint64 = 1024

// 	mask := capacity - 1

// 	res1 := hash % capacity // Классика
// 	res2 := hash & mask     // Битовый фильтр (хэш в конце)
// 	res3 := mask & hash     // Тот самый код автора (маска в начале)

// 	fmt.Printf("Математический остаток: %d\n", res1)
// 	fmt.Printf("Битовый (вариант 1):    %d\n", res2)
// 	fmt.Printf("Битовый (вариант 2):    %d\n", res3)
// }

func TestPutAndGet(t *testing.T) {
	// Создаём таблицу с небольшой capacity, чтобы
	// создать коллизию

	t.Run("wrong capacity", func(t *testing.T) {
		_, err := NewHashTableWithCapacity[string, int](3)
		if err.Error() != "capacity should be a power of 2" {
			t.Error("capacity validation falied")
		}
	})

	ht, err := NewHashTableWithCapacity[string, int](2)
	if err != nil {
		t.Fatalf("error creating a hashtable, err = %v", err)
	}

	t.Run("Put new element", func(t *testing.T) {
		ht.Put("apple", 100)
		val, err := ht.Get("apple")
		if err != nil || val != 100 {
			t.Errorf("Expected 100, got %v (err: %v)", val, err)
		}
	})

	t.Run("Update existing element", func(t *testing.T) {
		ht.Put("apple", 200)
	})

	t.Run("Force collision", func(t *testing.T) {
		ht.Put("banana", 300)
		ht.Put("cherry", 400)

		valBanana, errBanana := ht.Get("banana")
		valCherry, errCherry := ht.Get("cherry")

		if valBanana != 300 || valCherry != 400 {
			t.Errorf("Collision handling failed!")
		}

		if errBanana != nil || errCherry != nil {
			t.Errorf("Collision handling failed!")
		}
	})
}

func TestContains(t *testing.T) {
	ht := NewHashTable[string, string]()
	ht.Put("user1", "Lyolya")

	t.Run("Exisitng key", func(t *testing.T) {
		if !ht.Contains("user1") {
			t.Errorf("Expected Contains to return true for 'user1'")
		}
	})

	t.Run("Missing key", func(t *testing.T) {
		if ht.Contains("user2") {
			t.Errorf("Expected Contains to return false for 'user2'")
		}
	})
}

func TestRemove(t *testing.T) {
	// Небольшая таблица
	ht, err := NewHashTableWithCapacity[int, string](2)
	if err != nil {
		t.Fatal("should not be creating error")
	}

	ht.Put(1, "A")
	ht.Put(2, "B")

	// Должна быть коллизия
	ht.Put(3, "C")

	t.Run("Remove existing", func(t *testing.T) {
		err := ht.Remove(2)
		if err != nil {
			t.Errorf("Failed to remove existing key: %v", err)
		}
		// Проверим успешность удаления
		if ht.Contains(2) {
			t.Errorf("Key 2 should have been deleted")
		}
	})

	t.Run("Remove not-existing", func(t *testing.T) {
		err := ht.Remove(99)
		if err == nil {
			t.Errorf("Expected error when removing non-existing key")
		}
	})
}
