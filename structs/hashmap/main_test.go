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
		val, err := ht.Get("apple")
		if err != nil {
			t.Errorf("Error should be nil, but err = %v", err)
		}
		if val != 200 {
			t.Errorf("val shoul be equal to 200, but val = %v", val)
		}
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

	t.Run("Get from empty bucket", func(t *testing.T) {
		hm, _ := NewHashTableWithCapacity[int, string](2)

		_, err := hm.Get(0)

		if err == nil || err.Error() != "key not found" {
			t.Errorf("expected 'key not found' error from empty bucket")
		}
	})

	t.Run("Not found in existing chin", func(t *testing.T) {
		hm, _ := NewHashTableWithCapacity[int, string](2)

		hm.Put(3, "hey")
		hm.Put(5, "hoop")
		hm.Put(7, "lalaley")

		// Ключ, который точно есть в таблице
		exitingKey := 3
		// Считаем его хеш
		targetIndex := hm.hash(exitingKey)

		// Ищем новый ключ, который даст такой же индекс
		var missingKey int
		for i := 100; i < 10000; i++ {
			if hm.hash(i) == targetIndex {
				missingKey = i
				break
			}
		}

		if missingKey == 0 {
			t.Fatalf("No collision found after 10000 attempts")
		}

		// Теперь мы точно уверены, missingKey пойдёт в непустую ячейку,
		// но самого ключа там нет
		_, err := hm.Get(missingKey)

		if err == nil || err.Error() != "key not found" {
			t.Errorf("expected 'key not found', got: %v", err)
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

func TestHashTable_EdgeCases(t *testing.T) {
	t.Run("Zero capacity", func(t *testing.T) {
		ht, err := NewHashTableWithCapacity[string, int](0)
		if err == nil {
			t.Fatal("expected error for zero capacity, got nil")
		}
		if err.Error() != "capacity cannot be zero" {
			t.Errorf("wrong error message: %v", err)
		}
		if ht != nil {
			t.Error("hashtable should be nil on error")
		}
	})
}

func TestKeysValues(t *testing.T) {
	ht, _ := NewHashTableWithCapacity[string, int](4)
	ht.Put("one", 1)
	ht.Put("two", 2)
	ht.Put("three", 3)

	t.Run("Keys and Values", func(t *testing.T) {
		keys := ht.Keys()
		values := ht.Values()

		if len(keys) != 3 {
			t.Errorf("expected 3 keys, got %d", len(keys))
		}
		if len(values) != 3 {
			t.Errorf("expected 3 values, got %d", len(values))
		}

		// Используем map для проверки соответствия
		keyMap := make(map[string]int)
		for _, k := range keys {
			keyMap[k], _ = ht.Get(k)
		}

		if keyMap["one"] != 1 || keyMap["two"] != 2 || keyMap["three"] != 3 {
			t.Errorf("missing some keys or values")
		}
	})
}

func TestClear(t *testing.T) {
	ht := NewHashTable[int, int]()
	ht.Clear()

	ht.Put(1, 2)
	ht.Put(3, 4)
	ht.Clear()

	if ht.size != 0 {
		t.Errorf("size should be 0 after Clear, got %d", ht.size)
	}

	if len(ht.Keys()) != 0 || len(ht.Values()) != 0 {
		t.Errorf("Keys() and Values() should be empty after Clear")
	}
}
