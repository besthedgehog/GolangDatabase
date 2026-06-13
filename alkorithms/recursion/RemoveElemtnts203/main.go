package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func removeElements(head *ListNode, val int) *ListNode {
// 	// Граничный случай
// 	if head == nil {
// 		return nil
// 	}
// 	// Шаг рекурсии
// 	// К текущей ноде привяжем очищенный хвост
// 	head.Next = removeElements(head.Next, val)

// 	if head.Val == val {
// 		return head.Next
// 	} else {
// 		return head
// 	}
// }

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy
	for current.Next != nil {
		// Перешагиваем через неподоходящий элемент
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}

func printUp(n int) {
	if n <= 0 {
		return
	}
	printUp(n - 1) // Сначала ныряем до самого дна!
	fmt.Println(n) // Действие ПОСЛЕ всплытия
}

// Вызов printUp(3) выведет: 1, 2, 3
