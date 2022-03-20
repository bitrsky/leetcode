package main

import "fmt"

/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 *
 * https://leetcode.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Hard (41.75%)
 * Likes:    3476
 * Dislikes: 390
 * Total Accepted:    262.3K
 * Total Submissions: 627.1K
 * Testcase Example:  '3\n3'
 *
 * The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
 *
 * By listing and labeling all of the permutations in order, we get the
 * following sequence for n = 3:
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * Given n and k, return the k^th permutation sequence.
 *
 *
 * Example 1:
 * Input: n = 3, k = 3
 * Output: "213"
 * Example 2:
 * Input: n = 4, k = 9
 * Output: "2314"
 * Example 3:
 * Input: n = 3, k = 1
 * Output: "123"
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 9
 * 1 <= k <= n!
 *
 *
 */

// @lc code=start
func getPermutation(n int, k int) string {
	f := make([]int, n)
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = i * f[i-1]
	}

	ans := make([]byte, n)
	valid := make([]int, n+1)
	for i := 0; i <= n; i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/f[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order <= 0 {
				ans[i-1] = byte(j) + '0'
				valid[j] = 0
				break
			}
		}
		k %= f[n-i]
	}
	return string(ans)
}

// @lc code=end
func main() {
	fmt.Println(getPermutation(5, 10))
}
