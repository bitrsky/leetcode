package main

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (46.38%)
 * Likes:    11070
 * Dislikes: 443
 * Total Accepted:    1.2M
 * Total Submissions: 2.5M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * You are given an array of k linked-lists lists, each linked-list is sorted
 * in ascending order.
 *
 * Merge all the linked-lists into one sorted linked-list and return it.
 *
 *
 * Example 1:
 *
 *
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * merging them into one sorted list:
 * 1->1->2->3->4->4->5->6
 *
 *
 * Example 2:
 *
 *
 * Input: lists = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: lists = [[]]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] is sorted in ascending order.
 * The sum of lists[i].length will not exceed 10^4.
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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}
	for len(lists) > 1 {
		nexts := []*ListNode{}
		for i := 0; i < len(lists)-1; i += 2 {
			merge := mergeTwoLists(lists[i], lists[i+1])
			nexts = append(nexts, merge)
		}
		if len(lists)%2 == 1 {
			nexts = append(nexts, lists[len(lists)-1])
		}
		lists = nexts
	}
	return lists[0]
}

// @lc code=end
func main() {

	lists := []*ListNode{}
	lists = append(lists, buildList([]int{1, 2, 3}))
	lists = append(lists, buildList([]int{4, 5, 6}))
	lists = append(lists, buildList([]int{7, 8, 9}))
	printList(mergeKLists(lists))
}
