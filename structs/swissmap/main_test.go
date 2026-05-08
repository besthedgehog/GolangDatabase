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

func TestSwissMap_Basic(t *testing.T) {
	m := NewSwissMap[string, int](2)

	m.Put("Go", 1)
	m.Put("Rust", 2)

	{
		val, ok := m.Get("Go")
		if !ok {
			t.Errorf("ok should be true, but %v", ok)
		}
		if val != 1 {
			t.Errorf("val should be equal to 1, but %v", val)
		}
	}

	{
		val, ok := m.Get("Rust")
		if !ok {
			t.Errorf("ok should be true, but %v", ok)
		}
		if val != 2 {
			t.Errorf("val should be equal to 2, but %v", val)
		}
	}

	{
		val, ok := m.Get("C")
		if ok {
			t.Errorf("ok should be false but %v", ok)
		}
		if val != 0 {
			t.Errorf("val should be 0, but %v", val)
		}
	}
}
