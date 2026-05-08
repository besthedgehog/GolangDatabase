package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// slow указывает на место, куда
	// мы положим слдующий подходящий элемент
	var slow int = 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			// Переносим в чистую зону
			nums[slow] = nums[fast]
			// Сдвигаем границу чистой зоны
			slow++
		}
	}
	// Возвращаем количество элементов чистой зоны
	return slow
}

// func removeElement(nums []int, val int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	var slow int = 0
// 	for fast := 0; fast < len(nums); fast++ {
// 		if nums[slow] > val {
// 			slow++
// 			continue
// 		}
// 		if nums[fast] > val {
// 			nums[fast] = nums[slow]
// 		}
// 	}
// 	return len(nums)
// }
