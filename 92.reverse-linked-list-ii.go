package main

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (43.14%)
 * Likes:    5770
 * Dislikes: 272
 * Total Accepted:    461.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Given the head of a singly linked list and two integers left and right where
 * left <= right, reverse the nodes of the list from position left to position
 * right, and return the reversed list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], left = 2, right = 4
 * Output: [1,4,3,2,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [5], left = 1, right = 1
 * Output: [5]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is n.
 * 1 <= n <= 500
 * -500 <= Node.val <= 500
 * 1 <= left <= right <= n
 *
 *
 *
 * Follow up: Could you do it in one pass?
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ph := &ListNode{Next: head}
	lp := ph
	for i := 1; lp != nil && i < left; lp, i = lp.Next, i+1 {
	}
	mh, cp := lp, lp.Next
	for i := left; cp != nil && cp.Next != nil && i < right; i = i + 1 {
		cn := cp.Next
		cp.Next, cn.Next, mh.Next = cn.Next, mh.Next, cn
	}
	return ph.Next
}

// @lc code=end
func main() {
	h := buildList([]int{1, 2, 3, 4, 5})
	printList(reverseBetween(h, 2, 4))
}
