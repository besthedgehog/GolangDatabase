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
