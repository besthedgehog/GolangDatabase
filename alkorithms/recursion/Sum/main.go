package main

import "fmt"

func SumOfSlice(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	return arr[0] + SumOfSlice(arr[1:])
}

// SumByIndex с использованием индексов
// чтобы не создавались заново слайсы в ходе рекурсии
func SumByIndex(arr []int, index int) int {
	if len(arr) == index {
		return 0
	}
	return arr[index] + SumByIndex(arr, index+1)
}

func main() {
	var a = []int{1, 2, 3}
	fmt.Println(SumOfSlice(a))

	fmt.Println(SumByIndex(a, 0))
}
