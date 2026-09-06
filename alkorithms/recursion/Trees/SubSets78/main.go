package main

// Given an integer array nums of unique elements,
// return all possible subsets (the power set).

// The solution set must not contain duplicate subsets.
// Return the solution in any order.

// Example 1:

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Example 2:

// Input: nums = [0]
// Output: [[],[0]]

func subsets(nums []int) [][]int {
	lenNums := len(nums)
	var result [][]int

	var decideElement func(index int, tmpResult []int)
	tmpResult := []int{}

	decideElement = func(index int, tmpResult []int) {
		if index == lenNums {
			result = append(result, tmpResult)
			return
		}

		// Сюда добавим новый элемент
		newArray := append([]int{}, tmpResult...)
		newArray = append(newArray, nums[index])
		decideElement(index+1, newArray)

		// А сюда не добавим
		decideElement(index+1, tmpResult)
	}
	decideElement(0, tmpResult)
	return result
}
