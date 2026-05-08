package main

import (
	"fmt"
	"testing"
)

func TestCreateListNodesFromSlice(t *testing.T) {
	head := CreateListNodesFromSlice([]int{1, 2, 3})
	fmt.Println(head.Val)
	fmt.Println(head.Next.Val)
	fmt.Println(head.Next.Next.Val)
}
