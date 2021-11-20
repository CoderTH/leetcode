package main

import "sort"

/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
//1.暴力：排序，然后找到所有的和谐子序列最大的长度
func findLHS(nums []int) (ans int) {
	sort.Ints(nums)
	begin := 0
	for end, num := range nums {
		for num-nums[begin] > 1 {
			begin++
		}
		if num-nums[begin] == 1 && end-begin+1 > ans {
			ans = end - begin + 1
		}
	}
	return
}

// @lc code=end
