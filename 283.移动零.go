package main

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes1(nums []int) {
	//1. 使用格外数组，遇到不是零的插入新的数组，遇到零跳过，最后再补零
	newNums := make([]int, len(nums))
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			newNums[index] = nums[i]
			index++
		}
	}
	//默认位置为0，所以不用手动补零
	copy(nums, newNums)
}
func moveZeroes(nums []int) {
	//2. 直接在原数组中操作，使用快慢指针
	pointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pointer] = nums[i]
			if i != pointer {
				nums[i] = 0
			}
			pointer++
		}
	}

}

//1. 使用格外数组，遇到不是零的插入新的数组，遇到零跳过，最后再补零
//2. 直接在原数组中操作，使用快慢指针
// @lc code=end
