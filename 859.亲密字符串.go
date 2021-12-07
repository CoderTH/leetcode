package main

/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		dic := make(map[byte]bool)
		for i := 0; i < len(s); i++ {
			if dic[s[i]] {
				return true
			}
			dic[s[i]] = true
		}
		return false
	}
	index := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if len(index) >= 2 {
				return false
			}
			index = append(index, i)
		}
	}
	if len(index) != 2 {
		return false
	}
	return s[index[0]] == goal[index[1]] && s[index[1]] == goal[index[0]]
}

// @lc code=end
