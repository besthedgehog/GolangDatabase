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
