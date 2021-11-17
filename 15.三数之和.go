package main

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start

//1. 暴力解法：排序后，把所有的三数之和都列出来,注意每一层重复问题 O(n^3)，会超时
//2. 三指针：排序后,固定一个指针，剩下两个指针不断向中间移动
func threeSum1(nums []int) [][]int {
	//1. 暴力解法：把所有的三数之和都列出来 O(n^3)
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				tamp := nums[i] + nums[j] + nums[k]
				if tamp == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	//2. 三指针：排序后,固定一个指针，剩下两个指针不断向中间移动
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[right] + nums[left]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// @lc code=end
