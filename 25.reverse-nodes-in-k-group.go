package main

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (50.03%)
 * Likes:    6239
 * Dislikes: 471
 * Total Accepted:    469.9K
 * Total Submissions: 938.9K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, reverse the nodes of the list k at a time,
 * and return the modified list.
 *
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes, in the end, should remain as it is.
 *
 * You may not alter the values in the list's nodes, only nodes themselves may
 * be changed.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [2,1,4,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,3,4,5], k = 3
 * Output: [3,2,1,4,5]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is n.
 * 1 <= k <= n <= 5000
 * 0 <= Node.val <= 1000
 *
 *
 *
 * Follow-up: Can you solve the problem in O(1) extra memory space?
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func len(head *ListNode) int {
	length, ln := 0, head
	for ; ln != nil; length, ln = length+1, ln.Next {
	}
	return length
}

func reverseK(head *ListNode, k int) (*ListNode, *ListNode) {
	if len(head) < k {
		return head, nil
	}
	start := &ListNode{Next: head}
	end := start.Next
	for i := 1; i < k; i++ {
		// 先把这个节点摘出来
		successor := end.Next
		end.Next = successor.Next

		// put the node in head
		tmp := start.Next
		start.Next = successor
		successor.Next = tmp

		if successor == nil {
			break
		}
	}
	return start.Next, end
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	head = &ListNode{Next: head}
	l := head
	for l != nil {
		start, end := reverseK(l.Next, k)
		l.Next = start
		l = end
	}
	return head.Next
}

// @lc code=end

func main() {
	head := reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 3)
	printList(head)
}
