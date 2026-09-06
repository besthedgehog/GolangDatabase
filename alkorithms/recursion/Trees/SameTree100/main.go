package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	var left bool = isSameTree(p.Left, q.Left)
	var right bool = isSameTree(p.Right, q.Right)
	return left && right
}

//	 B
//  / \
// A  C

func CreateSameTrees() (*TreeNode, *TreeNode) {
	A := &TreeNode{1, nil, nil}
	C := &TreeNode{3, nil, nil}
	B := &TreeNode{2, A, C}

	A1 := &TreeNode{1, nil, nil}
	C1 := &TreeNode{3, nil, nil}
	B1 := &TreeNode{2, A1, C1}
	return B, B1
}

func CreateDifferentTrees() (*TreeNode, *TreeNode) {
	A := &TreeNode{1, nil, nil}
	C := &TreeNode{3, nil, nil}
	B := &TreeNode{2, A, C}

	A1 := &TreeNode{1, nil, nil}
	C1 := &TreeNode{7, nil, nil}
	B1 := &TreeNode{9, A1, C1}
	return B, B1
}

func main() {
	B, B1 := CreateSameTrees()
	answer := isSameTree(B, B1)
	fmt.Println(answer)
	fmt.Println()
	A, A1 := CreateDifferentTrees()
	answer1 := isSameTree(A, A1)
	fmt.Println(answer1)
}
