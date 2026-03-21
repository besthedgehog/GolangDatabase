package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	var data int = 7
	node := newNode(data)
	fmt.Println(node)

	assert.Equal(t, 7, node.data)
}

func TestNewCyclicLinkedList(t *testing.T) {
	newList := NewCyclicLinkedList[int]()

	assert.Equal(t, 0, newList.length)
	assert.Nil(t, newList.head)
}

func TestSize(t *testing.T) {
	cycledList := NewCyclicLinkedList[int]()

	expSize := 0
	size := cycledList.Size()
	assert.Equal(t, expSize, size)

	cycledList.Add(7)
	expSize = 1
	size = cycledList.Size()
	assert.Equal(t, expSize, size)
}

func TestEmpty(t *testing.T) {
	cycledList := NewCyclicLinkedList[int]()

	assert.True(t, cycledList.IsEmpty())

	cycledList.Add(7)

	assert.False(t, cycledList.IsEmpty())
}

func TestForEach(t *testing.T) {
	cycledList := NewCyclicLinkedList[int]()

	var called bool = false
	cycledList.ForEach(func(data int) {
		called = true
	})

	assert.False(t, called, "should not call fn on empty list")

	for i := range 3 {
		cycledList.Add(i)
	}

	resultList := []int{}
	cycledList.ForEach(func(data int) {
		resultList = append(resultList, data)
	})

	// fmt.Println(resultList)

	assert.Equal(t, resultList, []int{2, 1, 0})
}

/////

func TestReverseForEach(t *testing.T) {
	cycledList := NewCyclicLinkedList[string]()

	var called bool = false

	cycledList.ReverseForEach(func(data string) {
		called = true
	})

	assert.False(t, called, "should not call fn on empty list")

	cycledList.Add("A")
	cycledList.Add("B")
	cycledList.Add("C")

	resultList := []string{}
	cycledList.ReverseForEach(func(data string) {
		resultList = append(resultList, data)
	})

	assert.Equal(t, []string{"C", "A", "B"}, resultList)
}

func TestPrintList(t *testing.T) {
	cycledList := NewCyclicLinkedList[int]()

	cycledList.Add(1)
	cycledList.Add(2)
	cycledList.Add(3)

	expectedList := []int{3, 2, 1}
	actualList := cycledList.PrintList()
	assert.Equal(t, expectedList, actualList)
}

func TestRotate(t *testing.T) {

	t.Run("Rotate for 360", func(t *testing.T) {

		// create new cycledList
		cycledList := NewCyclicLinkedList[int]()

		// fullfill it
		for i := range 3 {
			cycledList.Add(i)
		}

		expectedList := []int{2, 1, 0}
		realList := cycledList.PrintList()
		assert.Equal(t, expectedList, realList)
		assert.Equal(t, 2, cycledList.head.data)

		// change nothing
		cycledList.Rotate(0)
		assert.Equal(t, expectedList, realList)
		assert.Equal(t, 2, cycledList.head.data)

		// change nothing
		cycledList.Rotate(3)
		assert.Equal(t, expectedList, realList)
		assert.Equal(t, 2, cycledList.head.data)

	})

	t.Run("Rotate for 1", func(t *testing.T) {
		// create new cycledList
		cycledList := NewCyclicLinkedList[int]()

		// fullfill it
		for i := range 3 {
			cycledList.Add(i)
		}

		cycledList.Rotate(1)
		expectedList := []int{1, 0, 2}
		realList := cycledList.PrintList()
		assert.Equal(t, 1, cycledList.head.data)
		assert.Equal(t, expectedList, realList)

		cycledList.Rotate(2)
		expectedList = []int{2, 1, 0}
		realList = cycledList.PrintList()
		assert.Equal(t, 2, cycledList.head.data)
		assert.Equal(t, expectedList, realList)
	})

	t.Run("rotate for -1", func(t *testing.T) {
		// create new cycledList
		cycledList := NewCyclicLinkedList[int]()

		// fullfill it
		for i := range 3 {
			cycledList.Add(i)
		}

		expectedList := []int{2, 1, 0}
		realList := cycledList.PrintList()
		assert.Equal(t, 2, cycledList.head.data)
		assert.Equal(t, expectedList, realList)

		cycledList.Rotate(-1)

		expectedList = []int{0, 2, 1}
		realList = cycledList.PrintList()
		assert.Equal(t, 0, cycledList.head.data)
		assert.Equal(t, expectedList, realList)

	})

}

func TestAdd(t *testing.T) {
	// 1. Тест пустого списка
	cycledList := NewCyclicLinkedList[int]()
	val, err := cycledList.Value()

	assert.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	assert.Equal(t, 0, val) // 0 - это *new(int)

	// 2. Тест списка с элементами
	cycledList.Add(10)
	val, err = cycledList.Value()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	cycledList.Add(20)

	// Последним добавили 20, значит 20 стало новой головой (согласно твоей логике Add)
	val, err = cycledList.Value()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
}

func TestRemove(t *testing.T) {
	t.Run("Remove from empty list", func(t *testing.T) {
		cycledList := NewCyclicLinkedList[int]()
		ok := cycledList.Remove()

		assert.False(t, ok)
		assert.Equal(t, 0, cycledList.Size())
	})

	t.Run("Remove single element", func(t *testing.T) {
		cycledList := NewCyclicLinkedList[int]()
		cycledList.Add(5)

		ok := cycledList.Remove()

		assert.True(t, ok)
		assert.Equal(t, 0, cycledList.Size())
		assert.Nil(t, cycledList.head)
		assert.True(t, cycledList.IsEmpty())
	})

	t.Run("Remove from multiple elements", func(t *testing.T) {
		cycledList := NewCyclicLinkedList[int]()
		// Порядок добавления: 1, 2, 3.
		// Твой Add вставляет элементы так, что 3 становится головой.
		cycledList.Add(1)
		cycledList.Add(2)
		cycledList.Add(3)

		// Ожидаемое состояние до удаления: Head -> 3. (3 <-> 2 <-> 1 <-> 3)
		assert.Equal(t, 3, cycledList.Size())
		val, _ := cycledList.Value()
		assert.Equal(t, 3, val)

		// Удаляем голову (должна удалиться 3)
		ok := cycledList.Remove()
		assert.True(t, ok)
		assert.Equal(t, 2, cycledList.Size())

		// Новой головой должен стать элемент, который шел за 3.
		// Судя по твоей функции Add, за 3 идет 2.
		val, _ = cycledList.Value()
		assert.Equal(t, 2, val)

		// Проверяем, что кольцо не разорвалось.
		// Если голова 2, то следующий должен быть 1.
		assert.Equal(t, 1, cycledList.head.nextPtr.data)
		// А предыдущий перед 2 (он же конец кольца) должен быть 1.
		assert.Equal(t, 1, cycledList.head.prevPtr.data)
	})
}

func TestRemoveAll(t *testing.T) {
	cycledList := NewCyclicLinkedList[int]()

	// Удаление из пустого должно пройти без паник
	ok := cycledList.RemoveAll()
	assert.True(t, ok)
	assert.Equal(t, 0, cycledList.Size())

	// Наполняем и удаляем
	cycledList.Add(10)
	cycledList.Add(20)
	cycledList.Add(30)

	assert.Equal(t, 3, cycledList.Size())

	ok = cycledList.RemoveAll()

	assert.True(t, ok)
	assert.Equal(t, 0, cycledList.Size())
	assert.Nil(t, cycledList.head)
	assert.True(t, cycledList.IsEmpty())
}
