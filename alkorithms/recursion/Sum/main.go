package main

import "fmt"

func SumOfSlice(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	return arr[0] + SumOfSlice(arr[1:])
}

func main() {
	var a = []int{1, 2, 3}
	fmt.Println(SumOfSlice(a))
}
