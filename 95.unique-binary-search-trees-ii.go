/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (48.88%)
 * Likes:    4529
 * Dislikes: 295
 * Total Accepted:    297.1K
 * Total Submissions: 606.8K
 * Testcase Example:  '3'
 *
 * Given an integer n, return all the structurally unique BST's (binary search
 * trees), which has exactly n nodes of unique values from 1 to n. Return the
 * answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3
 * Output:
 * [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: [[1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generate(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	if l == r {
		return []*TreeNode{&TreeNode{Val: l}}
	}
	res := []*TreeNode{}
	for i := l; i <= r; i++ {
		for _, lTree := range generate(l, i-1) {
			for _, rTree := range generate(i+1, r) {
				res = append(res, &TreeNode{Val: i, Left: lTree, Right: rTree})
			}
		}
	}
	return res
}
func generateTrees(n int) []*TreeNode {
	ans := []*TreeNode{}
	for i := 1; i <= n; i++ {
		for _, lTree := range generate(1, i-1) {
			for _, rTree := range generate(i+1, n) {
				ans = append(ans, &TreeNode{Val: i, Left: lTree, Right: rTree})
			}
		}
	}
	return ans
}

// @lc code=end
