package main

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (48.53%)
 * Likes:    4477
 * Dislikes: 181
 * Total Accepted:    818K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,1,2]'
 *
 * Given the head of a sorted linked list, delete all duplicates such that each
 * element appears only once. Return the linked list sorted as well.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,1,2]
 * Output: [1,2]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,1,2,3,3]
 * Output: [1,2,3]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 300].
 * -100 <= Node.val <= 100
 * The list is guaranteed to be sorted in ascending order.
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
func deleteDuplicates(head *ListNode) *ListNode {
	head = &ListNode{Next: head}
	for l, r := head, head.Next; r != nil; r = r.Next {
		if r.Next != nil && r.Val == r.Next.Val {
			for ; r.Next != nil && r.Val == r.Next.Val; r = r.Next {
			}
			if r.Next == nil {
				l.Next = nil
			}
			continue
		}
		l.Next = r
		l = l.Next
	}
	return head.Next
}

// @lc code=end
func main() {
	head := buildList([]int{1, 1})
	head = deleteDuplicates(head)
	printList(head)

}
