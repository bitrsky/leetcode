package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (30.54%)
 * Likes:    15934
 * Dislikes: 1528
 * Total Accepted:    1.8M
 * Total Submissions: 5.8M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an integer array nums, return all the triplets [nums[i], nums[j],
 * nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] +
 * nums[k] == 0.
 *
 * Notice that the solution set must not contain duplicate triplets.
 *
 *
 * Example 1:
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Example 2:
 * Input: nums = []
 * Output: []
 * Example 3:
 * Input: nums = [0]
 * Output: []
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
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
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for _, two := range twoSum(nums[i+1:], -1*nums[i]) {
			ans = append(ans, []int{nums[i], two[0], two[1]})
		}
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{1}))
}
