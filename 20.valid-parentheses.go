package main

import "fmt"

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (40.61%)
 * Likes:    11190
 * Dislikes: 472
 * Total Accepted:    2M
 * Total Submissions: 4.8M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: s = "(]"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 *
 *
 */

// @lc code=start
func isMatch(l, r uint8) int {
	if (l == '(' && r == ')') || (l == '{' && r == '}') || (l == '[' && r == ']') {
		return 0
	}
	if (l == '(' || l == '{' || l == '[') && (r == '(' || r == '{' || r == '[') {
		return 1
	}
	return -1
}
func isValid(s string) bool {
	stack := []uint8{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}
		match := isMatch(stack[len(stack)-1], s[i])
		if match == -1 {
			return false
		}
		if match == 1 {
			stack = append(stack, s[i])
			continue
		}
		stack = stack[0 : len(stack)-1]
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

// @lc code=end
func main() {
	fmt.Println(isValid("({})"))
}
