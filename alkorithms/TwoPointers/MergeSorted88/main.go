package main

import "fmt"

// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Указатели на последние элементы массива
	p1, p2 := m-1, n-1
	// Указатель на конец итогового массива (куда будем записывать)
	pMerge := m + n - 1
	// Переносим всё из nums2, значит,
	// если p2 ушёл в минус, то мы всё перенесли
	for p2 >= 0 {
		// Если в nums1 ещё остались элементы и текущий элемент nums1 больше
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[pMerge] = nums1[p1]
			p1--
		} else {
			nums1[pMerge] = nums2[p2]
			p2--
		}
		// Сдвигаем указатель записи на шаг влево
		pMerge--
	}
}

func main() {
	// Output: [1,2,2,3,5,6]
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
