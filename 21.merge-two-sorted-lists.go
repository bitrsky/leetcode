package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (59.28%)
 * Likes:    10521
 * Dislikes: 986
 * Total Accepted:    1.9M
 * Total Submissions: 3.3M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * You are given the heads of two sorted linked lists list1 and list2.
 *
 * Merge the two lists in a one sorted list. The list should be made by
 * splicing together the nodes of the first two lists.
 *
 * Return the head of the merged linked list.
 *
 *
 * Example 1:
 *
 *
 * Input: list1 = [1,2,4], list2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 *
 *
 * Example 2:
 *
 *
 * Input: list1 = [], list2 = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: list1 = [], list2 = [0]
 * Output: [0]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both list1 and list2 are sorted in non-decreasing order.
 *
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	l := head
	for ; list1 != nil && list2 != nil; l = l.Next {
		if list1.Val < list2.Val {
			l.Next = list1
			list1 = list1.Next
			continue
		}
		l.Next = list2
		list2 = list2.Next
	}
	if list1 != nil {
		l.Next = list1
	}
	if list2 != nil {
		l.Next = list2
	}
	return head.Next
}

// @lc code=end

func main() {
	printList(mergeTwoLists(buildList([]int{1, 2, 3}), buildList([]int{})))
}
