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
