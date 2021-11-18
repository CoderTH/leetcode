package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start

//思路：
//第一阶：一种方法
//第二阶：两种方法
//第三阶：因为一次只能跨一步或者两部，所以第三阶
//只能是由第一阶或者第二阶爬上来的，故：1+2=3
//第n阶段：f(n) = f(n-1)+f(n-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// @lc code=end
