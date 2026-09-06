package main

// 102. Binary Tree Level Order Traversal
// Medium

// Hint
// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e.,
// from left to right, level by level).

// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Example 3:
// Input: root = []
// Output: []

// Не работает
// func levelOrder(root *TreeNode) [][]int {
// 	var result [][]int = [][]int{}

// 	var trace func(root *TreeNode, currentLevel []int)
// 	trace = func(root *TreeNode, currentLevel []int) {
// 		if root == nil {
// 			return
// 		}
// 		currentLevel = append(currentLevel, root.Val)
// 		result = append(result, currentLevel)
// 		nextList := []int{}
// 		trace(root.Left, nextList)
// 		trace(root.Right, nextList)
// 	}
// 	trace(root, []int{})
// 	return result
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	currentLevel := root

	for currentLevel != nil {
		currentValues := []int{}
		nextLevel := []int{}

		for _, node := range currentLevel
	}

}

func main()
