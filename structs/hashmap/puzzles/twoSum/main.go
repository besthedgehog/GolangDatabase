package main

import "fmt"

func LexPrintf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := hashmap[complement]; ok {
			return []int{i, j}
		}
		hashmap[num] = i
	}
	return make([]int, 0)
}
