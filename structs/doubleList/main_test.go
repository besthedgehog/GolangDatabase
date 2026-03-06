package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushTail(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

	list.PushTail(10)
	expected := 10
	assert.Equal(t, list.head.data, expected)
	assert.Equal(t, list.tail.data, expected)

	list.PushTail(17)
	assert.Equal(t, list.tail.data, 17)
	assert.Equal(t, list.head.data, 10)
}

func TestPushHead(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)

	list.PushHead(10)
	expected := 10
	assert.Equal(t, list.head.data, expected)
	assert.Equal(t, list.tail.data, expected)

	list.PushHead(17)
	assert.Equal(t, list.head.data, 17)
	assert.Equal(t, list.tail.data, 10)
}

func TestInsert(t *testing.T) {
	list := NewDoublyLinkedList[string]()

	list.PushTail("A")
	list.PushTail("B")
	assert.Equal(t, list.head.data, "A")
	assert.Equal(t, list.tail.data, "B")

	list.Insert(1, "X")

	assert.Equal(t, list.head.data, "A")
	assert.Equal(t, list.head.nextPtr.data, "X")
	assert.Equal(t, list.head.nextPtr.nextPtr.data, "B")
}

func TestGet(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	for i := range 3 {
		list.PushTail(i)
	}

	{
		value, err := list.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, 1, value, 1)
	}

	{
		value, err := list.Get(87)
		assert.Equal(t, 0, value)
		assert.ErrorContains(t, err, "index out of range")
	}
}

func TestRemove(t *testing.T) {
	list := NewDoublyLinkedList[string]()

	for _, i := range []string{"A", "B", "C", "D"} {
		list.PushTail(i)
	}

	assert.Equal(t, list.head.data, "A")
	assert.Equal(t, list.head.nextPtr.data, "B")
	assert.Equal(t, list.head.nextPtr.nextPtr.data, "C")
	assert.Equal(t, list.head.nextPtr.nextPtr.nextPtr.data, "D")

	value, _ := list.Get(2)
	assert.Equal(t, value, "C")

	list.Remove(2)
	assert.Equal(t, list.length, 3)

	assert.Equal(t, list.head.data, "A")
	assert.Equal(t, list.head.nextPtr.data, "B")
	assert.Equal(t, list.head.nextPtr.nextPtr.data, "D")

}

func TestIsEmpyy(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	if list.IsEmpty() != true {
		t.Error("Expected list to be empty")
	}

	list.PushTail(0)

	if list.IsEmpty() != false {
		t.Error("Expected list to not be empty")
	}

}

func TestForEach(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	var called bool = false
	list.ForEach(func(data int) {
		called = true
	})

	assert.False(t, called, "should not call fn on empty list")

	for i := range 3 {
		list.PushTail(i)
	}

	resultList := []int{}
	list.ForEach(func(data int) {
		resultList = append(resultList, data)
	})

	assert.Equal(t, resultList, []int{0, 1, 2}, "should call fn for each element")

}

func TestForEachReverse(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	var called bool = false
	list.ForEach(func(data int) {
		called = true
	})

	assert.False(t, called, "should not call fn on empty list")

	for i := range 3 {
		list.PushTail(i)
	}

	resultList := []int{}
	list.ForEachReverse(func(data int) {
		resultList = append(resultList, data)
	})

	assert.Equal(t, resultList, []int{2, 1, 0}, "should call fn for each element")

}
