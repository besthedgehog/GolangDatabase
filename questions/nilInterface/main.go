package main

import "fmt"

type MyError struct {
	msg string
}

// Почему мы создём метод именно у указателя?
func (e *MyError) Error() string {
	return e.msg
}

func getError() error {
	// Почему указатель на структуру?
	var myErr *MyError = nil
	return myErr
}

func main() {
	err := getError()
	// Почему false?
	fmt.Println(err == nil)
	fmt.Println()
	var a *MyError = nil
	fmt.Println(a == nil)
	b := &MyError{}
	fmt.Println(b == nil)
}
