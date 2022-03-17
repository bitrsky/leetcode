/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (57.99%)
 * Likes:    6864
 * Dislikes: 273
 * Total Accepted:    466.1K
 * Total Submissions: 803.1K
 * Testcase Example:  '3'
 *
 * Given an integer n, return the number of structurally unique BST's (binary
 * search trees) which has exactly n nodes of unique values from 1 to n.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 19
 *
 * G(n) = G(0)*G(n-1) + G(1)*G(n-2)+.......+ G(n-2)*G(1) + G(n-1)*G(0)
 */

// @lc code=start
func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0], nums[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}

// @lc code=end
