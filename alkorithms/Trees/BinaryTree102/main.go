package main

// 102. Binary Tree Level Order Traversal

// Hint
// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		currentValues := []int{}
		nextLevel := []*TreeNode{}

		for _, node := range currentLevel {
			currentValues = append(currentValues, node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}

		}
		result = append(result, currentValues)
		currentLevel = nextLevel
	}
	return result
}
