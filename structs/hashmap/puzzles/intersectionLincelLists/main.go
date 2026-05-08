package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/// listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]

func CreateListNodesFromSlice(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	var currNode *ListNode = &ListNode{
		Val:  a[len(a)-1],
		Next: nil,
	}
	for i := (len(a) - 2); i >= 0; i-- {
		currNode = &ListNode{
			Val:  a[i],
			Next: currNode,
		}
	}
	return currNode
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	if headA == nil || headB == nil {
// 		return nil
// 	}
// 	// Список посещений
// 	visited := make(map[*ListNode]bool)
// 	for currA := headA; currA != nil; currA = currA.Next {
// 		visited[currA] = true
// 	}
// 	for currB := headB; currB != nil; currB = currB.Next {
// 		if visited[currB] == true {
// 			return currB
// 		}
// 	}
// 	return nil
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
}
