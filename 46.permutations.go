package main

import "fmt"

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (71.22%)
 * Likes:    9179
 * Dislikes: 173
 * Total Accepted:    1.1M
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3]'
 *
 * Given an array nums of distinct integers, return all the possible
 * permutations. You can return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * Example 2:
 * Input: nums = [0,1]
 * Output: [[0,1],[1,0]]
 * Example 3:
 * Input: nums = [1]
 * Output: [[1]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * All the integers of nums are unique.
 *
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 1 {
		ans = append(ans, []int{nums[0]})
		return ans
	}

	for _, p := range permute(nums[1:]) {
		for i := 0; i <= len(p); i++ {
			a := append([]int{}, p[:i]...)
			a = append(a, nums[0])
			a = append(a, p[i:]...)
			ans = append(ans, a)
		}
	}
	return ans
}

// @lc code=end

func main() {
	nums := []int{5, 4, 6, 2}
	fmt.Println(permute(nums))
}
