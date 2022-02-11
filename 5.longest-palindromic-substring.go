package main

import "fmt"

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (31.66%)
 * Likes:    15836
 * Dislikes: 933
 * Total Accepted:    1.7M
 * Total Submissions: 5.3M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, return the longest palindromic substring in s.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "babad"
 * Output: "bab"
 * Explanation: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "cbbd"
 * Output: "bb"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consist of only digits and English letters.
 *
 *
 */

/*
 * Node: do special treatment for lengths 0 and 1
 */
// @lc code=start
func palindrome(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
		if s[left] != s[right] {
			return left + 1, right - 1
		}
	}
	return left + 1, right - 1
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var pal string
	for i := 0; i < len(s); i++ {
		if i > 0 {
			left, right := palindrome(s, i, i)
			if right-left+1 > len(pal) {
				pal = s[left : right+1]
			}
		}
		if i < len(s)-1 && s[i] == s[i+1] {
			left, right := palindrome(s, i, i+1)
			if right-left+1 > len(pal) {
				pal = s[left : right+1]
			}
		}
	}
	return pal
}

// @lc code=end
func main() {
	fmt.Println(longestPalindrome("a"))
}
