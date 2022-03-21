package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (34.38%)
 * Likes:    3463
 * Dislikes: 659
 * Total Accepted:    375.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * There is an integer array nums sorted in non-decreasing order (not
 * necessarily with distinct values).
 *
 * Before being passed to your function, nums is rotated at an unknown pivot
 * index k (0 <= k < nums.length) such that the resulting array is [nums[k],
 * nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
 * For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and
 * become [4,5,6,6,7,0,1,2,4,4].
 *
 * Given the array nums after the rotation and an integer target, return true
 * if target is in nums, or false if it is not in nums.
 *
 * You must decrease the overall operation steps as much as possible.
 *
 *
 * Example 1:
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 * Example 2:
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * nums is guaranteed to be rotated at some pivot.
 * -10^4 <= target <= 10^4
 *
 *
 *
 * Follow up: This problem is similar to Search in Rotated Sorted Array, but
 * nums may contain duplicates. Would this affect the runtime complexity? How
 * and why?
 *
 */

// @lc code=start
func search(nums []int, target int) bool {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] && nums[mid] == nums[r] {
			l, r = l+1, r-1
			continue
		}
		if target < nums[mid] {
			// 0 - target - mid
			if target >= nums[0] {
				r = mid - 1
			} else if nums[mid] >= nums[0] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] >= nums[0] {
				l = mid + 1
			} else if target >= nums[0] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false

}

// @lc code=end
func main() {
	nums := []int{1, 0, 1, 1, 1}
	fmt.Println(search(nums, 0))
}
