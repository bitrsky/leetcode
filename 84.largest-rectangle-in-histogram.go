package main

import "fmt"

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 *
 * https://leetcode.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (40.28%)
 * Likes:    9381
 * Dislikes: 138
 * Total Accepted:    487.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * Given an array of integers heights representing the histogram's bar height
 * where the width of each bar is 1, return the area of the largest rectangle
 * in the histogram.
 *
 *
 * Example 1:
 *
 *
 * Input: heights = [2,1,5,6,2,3]
 * Output: 10
 * Explanation: The above is a histogram where width of each bar is 1.
 * The largest rectangle is shown in the red area, which has an area = 10
 * units.
 *
 *
 * Example 2:
 *
 *
 * Input: heights = [2,4]
 * Output: 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= heights.length <= 10^5
 * 0 <= heights[i] <= 10^4
 *
 *
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	n, maxA := len(heights), 0
	ls, rs, s := make([]int, n), make([]int, n), [][2]int{{-1, -1}}
	for i := 0; i < n; i++ {
		rs[i] = n
	}
	for i := 0; i < n; i++ {
		for ; heights[i] <= s[len(s)-1][0]; s = s[:len(s)-1] {
			rs[s[len(s)-1][1]] = i
		}
		ls[i] = s[len(s)-1][1]
		s = append(s, [2]int{heights[i], i})
	}
	for i := 0; i < n; i++ {
		area := heights[i] * (rs[i] - ls[i] - 1)
		if area > maxA {
			maxA = area
		}
	}
	return maxA
}

// @lc code=end
func main() {
	nums := []int{2, 4}
	fmt.Println(largestRectangleArea(nums))
}
