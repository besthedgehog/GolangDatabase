package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	current := dummyNode
	carry := 0 // Для переноса
	// Пока есть хотя бы одна фифра в l1 или l2 или перенос не равен 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// Перенос (17 / 10 = 1)
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return dummyNode.Next
}

func main() {

}
