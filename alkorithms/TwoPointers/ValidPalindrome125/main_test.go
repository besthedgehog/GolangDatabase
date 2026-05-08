package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
