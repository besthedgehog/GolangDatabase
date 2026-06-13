package main

// SOLUTION 1
// func isHappy(n int) bool {
// 	// Записываем числа, которые уже встречались
// 	seen := make(map[int]bool)
// 	for n != 1 && !seen[n] {
// 		seen[n] = true
// 		n = getNextSum(n)
// 	}
// 	return n == 1
// }

func getNextSum(n int) int {
	var sum int
	for n > 0 {
		digit := n % 10 // Последняя цифра
		sum += digit * digit
		n = n / 10
	}
	return sum
}

// func main() {
// 	var number int = 13
// 	getNextSum(number)
// 	_ = number
// }

// SOLUTION 2
// Алгоритм Флойда или Зайца и Черепахи
// func isHappy(n int) bool {
// 	slow := n
// 	fast := getNextSum(n)
// 	for fast != 1 && slow != fast {
// 		slow = getNextSum(slow)
// 		fast = getNextSum(getNextSum(fast))
// 	}
// 	return fast == 1
// }

// Solution 3
// Recursion
func isHappy(n int) bool {
	// Инициализируем map чтобы
	// отслеживать числа, которые уже встречались
	seen := make(map[int]bool)
	return isHappyHelper(n, seen)
}

func isHappyHelper(n int, seen map[int]bool) bool {
	if n == 1 {
		return true
	}
	if seen[n] == true {
		return false
	}
	seen[n] = true
	return isHappyHelper(getNextSum(n), seen)
}
