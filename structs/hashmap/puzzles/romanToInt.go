// LeetCode 13 Roman to integer
package main

func romanToInt(s string) int {
	// Создадим хеш-таблицу
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var total int

	var length int = len(s)

	for i := range length {
		currentValue := romanMap[s[i]]

		// Если мы не в конце строки и текущий символ меньше предыдущего
		if i < length-1 && currentValue < romanMap[s[i+1]] {
			total -= currentValue
		} else {
			total += currentValue
		}
	}
	return total
}
