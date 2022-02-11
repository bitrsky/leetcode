package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (47.01%)
 * Likes:    5219
 * Dislikes: 231
 * Total Accepted:    744.4K
 * Total Submissions: 1.6M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an integer array nums of length n and an integer target, find three
 * integers in nums such that the sum is closest to target.
 *
 * Return the sum of the three integers.
 *
 * You may assume that each input would have exactly one solution.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,2,1,-4], target = 1
 * Output: 2
 * Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,0,0], target = 1
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}
func twoSumClosest(nums []int, target int) int {
	minVal := nums[0] + nums[len(nums)-1]
	for l, r := 0, len(nums)-1; l < r; {
		sum := nums[l] + nums[r]
		if abs(minVal-target) > abs(sum-target) {
			minVal = sum
		}
		if sum > target {
			r--
		} else {
			l++
		}
	}
	return minVal
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minVal := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		twoSum := twoSumClosest(nums[i+1:], target-nums[i])
		if abs(minVal-target) > abs(twoSum+nums[i]-target) {
			minVal = twoSum + nums[i]
		}
	}
	return minVal
}

// @lc code=end

func main() {

	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	//fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println(threeSumClosest([]int{0, -4, 1, -5}, 0))
}
