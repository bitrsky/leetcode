package main

import "fmt"

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (53.37%)
 * Likes:    14290
 * Dislikes: 858
 * Total Accepted:    1.3M
 * Total Submissions: 2.4M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * You are given an integer array height of length n. There are n vertical
 * lines drawn such that the two endpoints of the i^th line are (i, 0) and (i,
 * height[i]).
 *
 * Find two lines that together with the x-axis form a container, such that the
 * container contains the most water.
 *
 * Return the maximum amount of water a container can store.
 *
 * Notice that you may not slant the container.
 *
 *
 * Example 1:
 *
 *
 * Input: height = [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * Explanation: The above vertical lines are represented by array
 * [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the
 * container can contain is 49.
 *
 *
 * Example 2:
 *
 *
 * Input: height = [1,1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4
 *
 *
 */

// @lc code=start

func min(m, n int) int {
	if m > n {
		return n
	}
	return m
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
func maxArea(height []int) int {
	maxArea := 0
	for left, right := 0, len(height)-1; left < right; {
		maxArea = max(maxArea, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			cur := height[left]
			for ; left < right && height[left] <= cur; left++ {
			}
		} else {
			cur := height[right]
			for ; left < right && height[right] <= cur; right-- {
			}
		}
	}
	return maxArea
}

// @lc code=end
func main() {
	// fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
