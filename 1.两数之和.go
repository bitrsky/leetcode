package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	num2Index := make(map[int]int)
	for index, num := range nums {
		element := target - num
		if indexAnother, exist := num2Index[element]; exist {
			return []int{index, indexAnother}
		}
		num2Index[num] = index
	}
	return []int{}
}

// @lc code=end
