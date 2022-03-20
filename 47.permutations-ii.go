package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (53.26%)
 * Likes:    4466
 * Dislikes: 90
 * Total Accepted:    569.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers, nums, that might contain duplicates, return
 * all possible unique permutations in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,2]
 * Output:
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	vis := make([]bool, len(nums))
	var backtrack func(res []int)
	backtrack = func(res []int) {
		if len(res) == len(nums) {
			ans = append(ans, append([]int{}, res...))
			return
		}
		for i, num := range nums {
			if vis[i] {
				continue
			}
			if i > 0 && num == nums[i-1] && vis[i-1] {
				continue
			}
			vis[i] = true
			res = append(res, num)
			backtrack(res)
			res = res[:len(res)-1]
			vis[i] = false
		}
	}
	backtrack([]int{})
	return ans
}

// @lc code=end
func main() {

	nums := []int{3, 3, 0, 3}
	ans := permuteUnique(nums)
	fmt.Println(ans)
}
