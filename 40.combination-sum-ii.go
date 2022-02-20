package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (51.99%)
 * Likes:    4535
 * Dislikes: 123
 * Total Accepted:    512.9K
 * Total Submissions: 986K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sum to target.
 *
 * Each number in candidates may only be used once in the combination.
 *
 * Note: The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [10,1,2,7,6,1,5], target = 8
 * Output:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5
 * Output:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */

// @lc code=start
func combination(candidates []int, target int) [][]int {
	ans := [][]int{}
	if len(candidates) <= 0 || candidates[0] > target {
		return ans
	}
	next := 1
	for ; next < len(candidates); next += 1 {
		if candidates[next] != candidates[next-1] {
			break
		}
	}
	c := []int{}
	for i := 1; i <= next; i++ {
		c = append(c, candidates[0])
		if target == i*candidates[0] {
			ans = append(ans, c)
			continue
		}
		for _, a := range combination(candidates[next:], target-i*candidates[0]) {
			a = append(a, c...)
			ans = append(ans, a)
		}
	}

	for _, a := range combination(candidates[next:], target) {
		ans = append(ans, a)
	}
	return ans
}
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combination(candidates, target)

}

// @lc code=end
func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	fmt.Println(combinationSum2(candidates, target))
}
