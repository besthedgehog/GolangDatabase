package main

func removeDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return []int{0}
		// return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
		continue
	}
	return nums
	// return slow + 1
}
