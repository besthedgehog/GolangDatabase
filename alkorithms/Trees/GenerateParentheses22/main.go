package main

import "fmt"

// Given n pairs of parentheses, write a function
// to generate all combinations of well-formed
// parentheses.

// Example 1:

// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// Example 2:

// Input: n = 1
// Output: ["()"]

func generateParenthesis(n int) []string {
	result := []string{}
	var backtrack func(current string, open int, close int)

	backtrack = func(current string, open int, close int) {
		if len(current) == n*2 {
			result = append(result, current)
			return
		}
		if open < n {
			backtrack(current+"(", open+1, close)
		}
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return result
}

func main() {
	a := generateParenthesis(3)
	fmt.Println(a)
}
