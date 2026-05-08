package main

import (
	"fmt"
	"testing"
)

// func TestIsAnagram(t *testing.T) {
// 	word1 := "Alex"
// 	word2 := "xelA"
// 	fmt.Println(isAnagram(word1, word2))
// }

// func TestSomething(t *testing.T) {
// 	word := "Alex"
// 	fmt.Printf("%q, %c\n\n", word[0], word[0])

// }

func TestIsAnagram(t *testing.T) {
	hashTable := make(map[int]int)

	val, ok := hashTable[7]
	fmt.Println("Here we go")
	fmt.Println(val)
	fmt.Println(ok)
}
