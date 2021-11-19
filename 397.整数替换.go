/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start
//1. 贪心：如果是偶数，那么直接/2，如果基数会有几种情况例如当前n=5
// 	* 给n+1 n=6，n/2=3，n-1=2，n/2=1 一共4步
//	* 给n-1 n=4,n/2=2,n/2=1,一共3步
//明显给n-1是最佳选择，那么可重复循环的规律是什么呢
// 其实就是无论给n+1还是-1，最后/2的结果是偶数就行了，其中一个特例是3需要特殊判断一下
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

