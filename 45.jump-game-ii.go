package main

import "fmt"

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 *
 * https://leetcode.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Medium (36.33%)
 * Likes:    7194
 * Dislikes: 267
 * Total Accepted:    541.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers nums, you are initially positioned
 * at the first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Your goal is to reach the last index in the minimum number of jumps.
 *
 * You can assume that you can always reach the last index.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2. Jump
 * 1 step from index 0 to 1, then 3 steps to the last index.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,3,0,1,4]
 * Output: 2
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 1000
 *
 *
 */

// @lc code=start

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	maxPosition, curEnd, step := 0, 0, 0
	for i := 0; i < len(nums)-1 && curEnd < len(nums)-1; i++ {
		if i+nums[i] > maxPosition {
			maxPosition = i + nums[i]
		}
		if i == curEnd {
			step++
			curEnd = maxPosition
		}
	}
	return step
}

// @lc code=end
func main() {
	nums := []int{2, 3, 1, 1, 5}
	fmt.Println(jump(nums))
}
