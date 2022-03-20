package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (44.56%)
 * Likes:    12637
 * Dislikes: 503
 * Total Accepted:    1.3M
 * Total Submissions: 3M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given an array of intervals where intervals[i] = [starti, endi], merge all
 * overlapping intervals, and return an array of the non-overlapping intervals
 * that cover all the intervals in the input.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */

// @lc code=start

type byFirst [][]int

func (s byFirst) Len() int           { return len(s) }
func (s byFirst) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byFirst) Less(i, j int) bool { return s[i][0] < s[j][0] }

func merge(intervals [][]int) [][]int {
	sort.Sort(byFirst(intervals))
	ans := [][]int{}
	start, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[0] <= end {
			if end < interval[1] {
				end = interval[1]
			}
			continue
		}
		ans = append(ans, []int{start, end})
		start, end = interval[0], interval[1]
	}
	ans = append(ans, []int{start, end})
	return ans
}

// @lc code=end
func main() {
	intervals := [][]int{{2, 3}}
	ans := merge(intervals)
	fmt.Println(ans)
}
