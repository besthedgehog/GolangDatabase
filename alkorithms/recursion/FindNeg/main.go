package main

import (
	"fmt"
	"strconv"
)

func FindNegative(arr []int) error {
	if len(arr) == 0 {
		return nil
	}
	if arr[0] <= 0 {
		number := strconv.Itoa(arr[0])
		return fmt.Errorf(number)
	}
	return FindNegative(arr[1:])

}

func main() {
	var a []int = []int{1, 2, 3, -5, 7}
	fmt.Println(FindNegative(a))
}
