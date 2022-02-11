package main

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (37.83%)
 * Likes:    16288
 * Dislikes: 3485
 * Total Accepted:    2.5M
 * Total Submissions: 6.6M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order, and each of their nodes
 * contains a single digit. Add the two numbers and return the sum as a linked
 * list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 *
 * Example 1:
 *
 *
 * Input: l1 = [2,4,3], l2 = [5,6,4]
 * Output: [7,0,8]
 * Explanation: 342 + 465 = 807.
 *
 *
 * Example 2:
 *
 *
 * Input: l1 = [0], l2 = [0]
 * Output: [0]
 *
 *
 * Example 3:
 *
 *
 * Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * Output: [8,9,9,9,0,0,0,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in each linked list is in the range [1, 100].
 * 0 <= Node.val <= 9
 * It is guaranteed that the list represents a number that does not have
 * leading zeros.
 *
 *
 */
/*
type ListNode struct {
	Val  int
	Next *ListNode
}
*/

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	var carry int = 0
	l := head
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		sum := l1.Val + l2.Val + carry
		l1.Val = sum % 10
		carry = sum / 10
		l.Next = l1
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}
	if carry == 0 {
		return head.Next
	}

	for ; l.Next != nil; l = l.Next {
		sum := l.Next.Val + carry
		l.Next.Val = sum % 10
		carry = sum / 10
	}

	if carry == 0 {
		return head.Next
	}
	l.Next = &ListNode{}
	l.Next.Val = carry
	return head.Next
}

// @lc code=end

func main() {
	l1 := buildList([]int{9, 9, 9})
	l2 := buildList([]int{1})
	printList(addTwoNumbers(l1, l2))
}
