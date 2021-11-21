package main

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */
type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var maxNode int
	for i := 0; i < len(root.Children); i++ {
		maxNode = max(maxNode, maxDepth(root.Children[i]))
	}
	return maxNode + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
