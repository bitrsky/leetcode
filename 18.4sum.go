package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (37.00%)
 * Likes:    5551
 * Dislikes: 634
 * Total Accepted:    540K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers, return an array of all the unique
 * quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
 *
 *
 * 0 <= a, b, c, d < n
 * a, b, c, and d are distinct.
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,0,-1,0,-2,2], target = 0
 * Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,2,2,2,2], target = 8
 * Output: [[2,2,2,2]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start
func stepUp(index int, nums []int) int {
	for ; index > 0 && index < len(nums) && nums[index] == nums[index-1]; index++ {
	}
	return index
}
func stepDown(index int, nums []int) int {
	for ; index >= 0 && index < len(nums)-1 && nums[index] == nums[index+1]; index-- {
	}
	return index
}
func twoSum(nums []int, target int) [][]int {
	ans := [][]int{}
	for left, right := 0, len(nums)-1; left < right; {
		sum := nums[left] + nums[right]
		if sum == target {
			ans = append(ans, []int{nums[left], nums[right]})
			left = stepUp(left+1, nums)
			right = stepDown(right-1, nums)
			continue
		}
		if sum < target {
			left = stepUp(left+1, nums)
		} else {
			right = stepDown(right-1, nums)
		}
	}
	return ans
}
func threeSum(nums []int, target int) [][]int {
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for _, two := range twoSum(nums[i+1:], target-nums[i]) {
			ans = append(ans, append([]int{nums[i]}, two...))
		}
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for _, three := range threeSum(nums[i+1:], target-nums[i]) {
			ans = append(ans, append([]int{nums[i]}, three...))
		}
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
