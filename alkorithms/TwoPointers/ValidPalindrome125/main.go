package main

import "unicode"

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		lRune := rune(s[left])
		rRune := rune(s[right])
		// skip left
		if !unicode.IsLetter(lRune) && !unicode.IsDigit(lRune) {
			left++
			continue
		}
		// skip right
		if !unicode.IsLetter(rRune) && !unicode.IsDigit(rRune) {
			right--
			continue
		}

		if unicode.ToLower(lRune) != unicode.ToLower(rRune) {
			return false
		}
		left++
		right--
	}
	return true
}
