package main

import (
	"fmt"
	"testing"
)

func TestBites(t *testing.T) {
	var a = 0b1000
	var b = a >> 2 // 1000 -> 0010
	var c = a << 2 // 1000 -> 100000

	// 1000 это 2**4 = 8
	fmt.Printf("a = %b or %v", a, a)
	fmt.Println()
	// 10 это 2**1 = 2
	fmt.Printf("b = %b or %v\n", b, b)
	fmt.Println()
	// 100000 это 2**5 = 32
	fmt.Printf("c = %b or %v\n", c, c)
}
