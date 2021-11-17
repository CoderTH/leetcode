package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//1. hash: 遍历一遍将所有的元素存在hash表中，每次进行查找，如果能找到说明形成环
//2. 快慢指针： 慢指针一次走一步，快指针两步，如果快指针==慢指针说明追到

func hasCycle1(head *ListNode) bool {
	dic := make(map[*ListNode]int, 0)
	for head != nil {
		if dic[head] != 0 {
			return true
		}
		dic[head]++
		head = head.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		if slow == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// @lc code=end
