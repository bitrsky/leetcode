package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (63.88%)
 * Likes:    9046
 * Dislikes: 207
 * Total Accepted:    961.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given an array of distinct integers candidates and a target integer target,
 * return a list of all unique combinations of candidates where the chosen
 * numbers sum to target. You may return the combinations in any order.
 *
 * The same number may be chosen from candidates an unlimited number of times.
 * Two combinations are unique if the frequency of at least one of the chosen
 * numbers is different.
 *
 * It is guaranteed that the number of unique combinations that sum up to
 * target is less than 150 combinations for the given input.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7
 * Output: [[2,2,3],[7]]
 * Explanation:
 * 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple
 * times.
 * 7 is a candidate, and 7 = 7.
 * These are the only two combinations.
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8
 * Output: [[2,2,2,2],[2,3,3],[3,5]]
 *
 *
 * Example 3:
 *
 *
 * Input: candidates = [2], target = 1
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * All elements of candidates are distinct.
 * 1 <= target <= 500
 *
 *
 */

// @lc code=start
func combination(candidates []int, target int) [][]int {
	ans := [][]int{}
	if len(candidates) <= 0 || candidates[0] > target {
		return ans
	}
	if candidates[0] == target {
		ans = append(ans, []int{candidates[0]})
		return ans
	}
	// 不选当前值
	for _, a := range combination(candidates[1:], target) {
		ans = append(ans, a)
	}

	// 选当前值
	for _, a := range combination(candidates, target-candidates[0]) {
		a = append(a, candidates[0])
		ans = append(ans, a)
	}
	return ans
}
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combination(candidates, target)
}

// @lc code=end
func main() {
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9
	fmt.Println(combinationSum(candidates, target))
}
