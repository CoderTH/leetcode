package main

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
//1. 暴力解法：直接遍历数组查找
//2. 双指针（和暴力差不多）
//3. 二分法
func search1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == target {
			return l
		} else if nums[r] == target {
			return r
		} else {
			l++
			r--
		}
	}
	return -1
}
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end
