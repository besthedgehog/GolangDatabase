// 141. Linked List Cycle

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	holderMap := make(map[*ListNode]bool)

	currentNode := head

	for {
		if currentNode == nil {
			return false
		}
		if holderMap[currentNode] == true {
			return true
		}
		holderMap[currentNode] = true
		currentNode = currentNode.Next
	}
}

func hasCycle2(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
