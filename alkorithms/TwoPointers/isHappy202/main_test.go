package main

import (
	"fmt"
	"testing"
)

func TestSomthing(t *testing.T) {
	// var number int = 123
	number := 123

	fmt.Println(number)
	fmt.Println(number % 10)
	fmt.Println(number / 10)
}

// func isHappy(n int) bool {
// 	slow := n
// 	fast := getNextSum(n)

// 	// Пока заяц не добежал до финиша (1)
// 	// и пока заяц не догнал черепаху (slow == fast)
// 	for fast != 1 && slow != fast {
// 		slow = getNextSum(slow)             // Черепаха делает 1 шаг
// 		fast = getNextSum(getNextSum(fast)) // Заяц делает 2 шага
// 	}

// 	// Если заяц наступил на 1, значит число счастливое
// 	return fast == 1
// }

// RECURSION
//
//
//
//
// func isHappy(n int) bool {
// 	// Инициализируем мапу при первом вызове
// 	seen := make(map[int]bool)
// 	return isHappyHelper(n, seen)
// }

// // Вспомогательная функция, которая делает всю рекурсивную работу
// func isHappyHelper(n int, seen map[int]bool) bool {
// 	if n == 1 {
// 		return true
// 	}
// 	if seen[n] {
// 		return false // Мы тут уже были, это бесконечный цикл!
// 	}

// 	// Запоминаем текущее число
// 	seen[n] = true

// 	// Шагаем дальше в рекурсию с новым числом
// 	next := getNextSum(n)
// 	return isHappyHelper(next, seen)
// }
