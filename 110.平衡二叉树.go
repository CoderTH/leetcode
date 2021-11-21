package main

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if _, ok := maxDepth(root); ok {
		return true
	}
	return false
}
func maxDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, leftOK := maxDepth(root.Left)
	right, rightOK := maxDepth(root.Right)
	if !leftOK || !rightOK || abs(left-right) > 1 {
		return 0, false
	}
	if left > right {
		return left + 1, true
	}
	return right + 1, true

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
