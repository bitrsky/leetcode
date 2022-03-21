package main

import "fmt"

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (62.71%)
 * Likes:    3892
 * Dislikes: 129
 * Total Accepted:    497.1K
 * Total Submissions: 784.8K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of the range [1, n].
 *
 * You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1, k = 1
 * Output: [[1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 20
 * 1 <= k <= n
 *
 *
 */

// @lc code=start
func combine(n int, k int) [][]int {

	ans := [][]int{}
	res := make([]int, 0, k)
	var backtrack func(n, x int)
	backtrack = func(m, x int) {
		if x == 0 && len(res) == k {
			ans = append(ans, append([]int{}, res...))
			return
		}
		if m > n {
			return
		}
		backtrack(m+1, x)

		res = append(res, m)
		backtrack(m+1, x-1)
		res = res[:len(res)-1]
	}
	backtrack(1, k)
	return ans
}

// @lc code=end
func main() {

	fmt.Println(combine(4, 2))

}
