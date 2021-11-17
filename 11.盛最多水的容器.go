package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
//1. 枚举：把柱子两两组合，算出其中最大的值
//2. 两个指针从外向内收敛
func maxArea1(height []int) int {
	//1. 枚举：把柱子两两组合，算出其中最大的值(会超时)
	ans := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			tamp := (j - i) * min(height[i], height[j])
			ans = max(ans, tamp)
		}
	}
	return ans
}

func maxArea(height []int) int {
	//2. 两个指针从外向内收敛
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		tamp := (right - left) * min(height[right], height[left])
		ans = max(ans, tamp)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
