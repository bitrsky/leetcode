package main

import "fmt"

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (31.04%)
 * Likes:    7051
 * Dislikes: 242
 * Total Accepted:    454.8K
 * Total Submissions: 1.5M
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()".
 *
 *
 * Example 2:
 *
 *
 * Input: s = ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()".
 *
 *
 * Example 3:
 *
 *
 * Input: s = ""
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 3 * 10^4
 * s[i] is '(', or ')'.
 *
 *
 */

// @lc code=start
func longestValidParentheses(s string) int {
	f := make([]int, len(s), len(s))
	for i := 0; i < len(s); i += 1 {
		// ( 直接入
		if s[i] == '(' {
			f[i] = i
			continue
		}
		// i == 0 直接入, f[i-1] 已经上完全匹配
		if i < 1 || f[i-1] < 0 {
			f[i] = i
			continue
		}

		// 没匹配上
		if s[f[i-1]] == ')' {
			f[i] = i
			continue
		}
		// f[i] 完全匹配
		if f[i-1]-1 < 0 {
			f[i] = -1
			continue
		}
		f[i] = f[f[i-1]-1]
	}
	max := 0
	for i := 0; i < len(f); i++ {
		if i-f[i] > max {
			max = i - f[i]
		}
	}
	return max

}

// @lc code=end
func main() {

	// s := "(()"
	s := "()"
	fmt.Println(longestValidParentheses(s))

}
