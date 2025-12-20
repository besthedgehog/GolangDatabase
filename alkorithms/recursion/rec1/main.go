// Напиши функцию, которая принимает целое неотрицательное
// число n и возвращает сумму всех чисел от 1 до n.

package main

import (
	"fmt"
	"os"
)

func sum(n int) int {
	if n < 0 {
		os.Exit(1)
	}
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

// Более строгая версия с проверкой на ошибки
func Sum(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n cannot be negative")
	}
	if n == 0 {
		return 0, nil
	}
	res, err := Sum(n - 1)
	if err != nil {
		return 0, err
	}
	return n + res, nil
}

func main() {
	n := 4
	fmt.Println(sum(n))
	fmt.Println(Sum(n + 1))
}
