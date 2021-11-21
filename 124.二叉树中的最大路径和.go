package main

import "math"

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
//dfs
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := max(dfs(root.Left), 0)
		rightSum := max(dfs(root.Right), 0)
		tamp := leftSum + rightSum + root.Val
		ans = max(ans, tamp)
		return root.Val + max(leftSum, rightSum)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
