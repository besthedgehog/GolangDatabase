package main

import (
	"fmt"
)

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along
// the longest path from the root node down to the farthest
// leaf node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var currentResult int = 0
	var fallDepth func(root *TreeNode, currentLevel int)

	fallDepth = func(root *TreeNode, currentLevel int) {
		if root == nil {
			return
		}
		if currentLevel > currentResult {
			currentResult = currentLevel
		}
		fallDepth(root.Left, currentLevel+1)
		fallDepth(root.Right, currentLevel+1)
	}

	// Запуск с единицы, так как считаем
	// количество нод
	fallDepth(root, 1)
	return currentResult
}

func CreateTree2() *TreeNode {
	B := &TreeNode{9, nil, nil}
	A := &TreeNode{10, B, nil}
	return A
}

func CreateTree3() *TreeNode {
	C := &TreeNode{10, nil, nil}
	B := &TreeNode{9, C, nil}
	A := &TreeNode{10, B, nil}
	return A
}

func CreateTree4() *TreeNode {
	D := &TreeNode{10, nil, nil}
	C := &TreeNode{10, D, nil}
	B := &TreeNode{9, C, nil}
	A := &TreeNode{10, B, nil}
	return A
}

func CreateTree33() *TreeNode {
	D := &TreeNode{5, nil, nil}
	E := &TreeNode{4, nil, nil}
	C := &TreeNode{3, nil, nil}
	B := &TreeNode{2, D, E}
	A := &TreeNode{1, B, C}
	return A
}

func main() {

	fmt.Println(maxDepth(CreateTree2()))
	fmt.Println(maxDepth(CreateTree3()))
	fmt.Println(maxDepth(CreateTree4()))

	{
		A := CreateTree33()
		result := maxDepth(A)
		fmt.Println(result)
	}
	fmt.Println(maxDepth(&TreeNode{}))
	fmt.Println(maxDepth(nil))
}

// Проще

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := 1 + maxDepth(root.Right)
	left := 1 + maxDepth(root.Left)
	return max(right, left)
}

// Или

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}
