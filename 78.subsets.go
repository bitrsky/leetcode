package main

import "fmt"

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (69.81%)
 * Likes:    9452
 * Dislikes: 151
 * Total Accepted:    1M
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3]'
 *
 * Given an integer array nums of unique elements, return all possible subsets
 * (the power set).
 *
 * The solution set must not contain duplicate subsets. Return the solution in
 * any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3]
 * Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0]
 * Output: [[],[0]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10
 * -10 <= nums[i] <= 10
 * All the numbers of nums are unique.
 *
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	n := len(nums)
	max := 1 << n
	ans := [][]int{}
	for i := 0; i < max; i++ {
		res := []int{}
		for j := 0; j < n; j++ {
			if i>>j%2 == 1 {
				res = append(res, nums[j])
			}
		}
		ans = append(ans, res)
	}
	return ans
}

// @lc code=end

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsets(nums))
}
