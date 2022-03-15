package main

/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 *
 * https://leetcode.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Medium (45.96%)
 * Likes:    4004
 * Dislikes: 149
 * Total Accepted:    272.8K
 * Total Submissions: 592.9K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * You are given the root of a binary search tree (BST), where the values of
 * exactly two nodes of the tree were swapped by mistake. Recover the tree
 * without changing its structure.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,3,null,null,2]
 * Output: [3,1,null,null,2]
 * Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3
 * makes the BST valid.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [3,1,4,null,null,2]
 * Output: [2,1,4,null,null,3]
 * Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2
 * and 3 makes the BST valid.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [2, 1000].
 * -2^31 <= Node.val <= 2^31 - 1
 *
 *
 *
 * Follow up: A solution using O(n) space is pretty straight-forward. Could you
 * devise a constant O(1) space solution?
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

func inOrderTraversal(cur *TreeNode) []*TreeNode {
	visit, stack := []*TreeNode{}, []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visit = append(visit, cur)
		cur = cur.Right
	}
	return visit
}

func findErrorNode(list []*TreeNode) (l *TreeNode, r *TreeNode) {

	for i := 0; i < len(list)-1; i++ {
		if list[i].Val <= list[i+1].Val {
			continue
		}
		if l == nil {
			l = list[i]
		}
		r = list[i+1]
	}
	return l, r
}
func recoverTree(root *TreeNode) {
	list := inOrderTraversal(root)
	l, r := findErrorNode(list)
	if l == nil && r == nil {
		return
	}
	l.Val, r.Val = r.Val, l.Val
}

// @lc code=end
func main() {
	nums := []int{1, 3, -1, -1, 2}
	root := buildTree(nums)
	recoverTree(root)
	printTree(root)
}
