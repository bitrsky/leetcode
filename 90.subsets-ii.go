package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (52.75%)
 * Likes:    4648
 * Dislikes: 144
 * Total Accepted:    477.8K
 * Total Submissions: 904.6K
 * Testcase Example:  '[1,2,2]'
 *
 * Given an integer array nums that may contain duplicates, return all possible
 * subsets (the power set).
 *
 * The solution set must not contain duplicate subsets. Return the solution in
 * any order.
 *
 *
 * Example 1:
 * Input: nums = [1,2,2]
 * Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
 * Example 2:
 * Input: nums = [0]
 * Output: [[],[0]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10
 * -10 <= nums[i] <= 10
 *
 *
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{[]int{}}
	sort.Ints(nums)
	n := len(nums)
	for m := 0; m < (1 << n); m++ {
		a, isOk := []int{}, true
		for i := 0; i < n; i++ {
			if m>>i&1 == 0 {
				continue
			}
			if i > 0 && m>>(i-1)&1 == 0 && nums[i] == nums[i-1] {
				isOk = false
				break
			}
			a = append(a, nums[i])
		}

		if isOk && len(a) > 0 {
			ans = append(ans, a)
		}
	}
	return ans
}

// @lc code=end

func main() {
	nums := []int{1, 2, 2}
	ans := subsetsWithDup(nums)
	fmt.Println(ans)
}
