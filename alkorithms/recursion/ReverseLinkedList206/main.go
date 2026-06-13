package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	if head.Next == nil {
// 		return head
// 	}

// 	newHead := reverseList(head.Next)

// 	// newHead = Node 2
// 	// head = Node 1

// 	//       +----------------+
// 	//       |                |
// 	//       v                |
// 	//    [ 1 ] ------------> [ 2 ]
// 	head.Next.Next = head
// 	head.Next = nil

// 	return newHead
// }

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		// 1. Прячем следующий вагон в карман, чтобы не потерять цепочку
		nextTemp := curr.Next
		// 2. РАЗВОРАЧИВАЕМ СТРЕЛКУ! Текущая нода теперь смотрит назад.
		curr.Next = prev
		// 3. Шагаем вперёд
		prev = curr
		curr = nextTemp
	}
	// Когда curr улетит в nil (за конец списка),
	// prev будет стоять ровно на последней ноде. Она и стала новой головой!
	return prev
}

func NewList() *ListNode {
	B := &ListNode{Val: 2, Next: nil}
	A := &ListNode{Val: 1, Next: B}
	return A
}

func main() {
	head := NewList()
	head = reverseList(head)
	fmt.Println(head)
}
