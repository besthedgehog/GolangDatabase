package main

import (
	"fmt"
	"os"
)

func Fib(n int) int {

	if n < 0 {
		os.Exit(1)
		fmt.Println("n cannot be negative")
	}

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

func main() {
	var a int = 7

	fmt.Println(Fib(a))

}
