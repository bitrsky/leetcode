package main

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (33.72%)
 * Likes:    4828
 * Dislikes: 1261
 * Total Accepted:    534.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, rotate the list to the right by k
 * places.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [4,5,1,2,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [0,1,2], k = 4
 * Output: [2,0,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 500].
 * -100 <= Node.val <= 100
 * 0 <= k <= 2 * 10^9
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	end, length := head, 1
	for ; end.Next != nil; end = end.Next {
		length++
	}

	k = length - k%length
	end.Next, head = head, end
	for i := 0; i < k; i++ {
		head = head.Next
	}
	end, head = head, head.Next
	end.Next = nil
	return head
}

// @lc code=end
func main() {

	head := buildList([]int{1, 2, 3, 4, 5})
	printList(head)
	head = rotateRight(head, 2)
	printList(head)
}
