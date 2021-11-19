/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start
func integerReplacement(n int) int {
	count := 0
	for n > 1 {
		if n&1 == 0 {
			n /= 2
		} else {
			if (n-1)/2&1 == 1 && n != 3 {
				n += 1
			} else {
				n -= 1
			}
		}
		count++
	}
	return count
}

// @lc code=end

