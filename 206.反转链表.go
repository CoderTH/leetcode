package main

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 题很简单，直接进行模拟反转就好
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		cur := head.Next
		head.Next = prev
		prev = head
		head = cur
	}
	//
	return prev
}

// @lc code=end
