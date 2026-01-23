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
