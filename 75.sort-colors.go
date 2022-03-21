package main

import "fmt"

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 *
 * https://leetcode.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (53.90%)
 * Likes:    9124
 * Dislikes: 397
 * Total Accepted:    953.8K
 * Total Submissions: 1.8M
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * Given an array nums with n objects colored red, white, or blue, sort them
 * in-place so that objects of the same color are adjacent, with the colors in
 * the order red, white, and blue.
 *
 * We will use the integers 0, 1, and 2 to represent the color red, white, and
 * blue, respectively.
 *
 * You must solve this problem without using the library's sort function.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,0,2,1,1,0]
 * Output: [0,0,1,1,2,2]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,0,1]
 * Output: [0,1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * n == nums.length
 * 1 <= n <= 300
 * nums[i] is either 0, 1, or 2.
 *
 *
 *
 * Follow up: Could you come up with a one-pass algorithm using only constant
 * extra space?
 *
 */

// @lc code=start
func sort(nums []int, v int) int {
	l, r := 0, len(nums)-1
	for l < r {
		for ; l < r && nums[l] <= v; l++ {
		}
		for ; l < r && nums[r] > v; r-- {
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return r

}
func sortColors(nums []int) {
	r := sort(nums, 1)
	sort(nums[:r+1], 0)
}

// @lc code=end
func main() {

	nums := []int{1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
