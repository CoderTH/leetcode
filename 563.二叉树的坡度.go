package main

/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func findTilt(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sumLeft := dfs(root.Left)
	sumRight := dfs(root.Right)
	res += abs(sumLeft - sumRight)
	return sumLeft + sumRight + root.Val
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
