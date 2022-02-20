package main

import "fmt"

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (55.59%)
 * Likes:    16680
 * Dislikes: 235
 * Total Accepted:    1M
 * Total Submissions: 1.8M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it can trap after raining.
 *
 *
 * Example 1:
 *
 *
 * Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 * Explanation: The above elevation map (black section) is represented by array
 * [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue
 * section) are being trapped.
 *
 *
 * Example 2:
 *
 *
 * Input: height = [4,2,0,3,2,5]
 * Output: 9
 *
 *
 *
 * Constraints:
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */

// @lc code=start
func trap(height []int) int {
	area := 0
	l, r := 0, len(height)-1
	for l < r {
		lh, rh := height[l], height[r]
		if lh < rh {
			for ; height[l] <= lh && l < r; l++ {
				area += lh - height[l]
			}
		} else {
			for ; height[r] <= rh && r > l; r-- {
				area += rh - height[r]
			}
		}
	}
	return area
}

// @lc code=end
func main() {
	nums := []int{4, 2, 3}
	fmt.Println(trap(nums))

}
