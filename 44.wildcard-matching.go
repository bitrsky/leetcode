package main

import "fmt"

/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 *
 * https://leetcode.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (26.43%)
 * Likes:    4403
 * Dislikes: 197
 * Total Accepted:    371.2K
 * Total Submissions: 1.4M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement wildcard pattern
 * matching with support for '?' and '*' where:
 *
 *
 * '?' Matches any single character.
 * '*' Matches any sequence of characters (including the empty sequence).
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 *
 * Example 1:
 *
 *
 * Input: s = "aa", p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "aa", p = "*"
 * Output: true
 * Explanation: '*' matches any sequence.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "cb", p = "?a"
 * Output: false
 * Explanation: '?' matches 'c', but the second letter is 'a', which does not
 * match 'b'.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length, p.length <= 2000
 * s contains only lowercase English letters.
 * p contains only lowercase English letters, '?' or '*'.
 *
 *
 */

// @lc code=start
func match(s, p byte) bool {
	if p == '?' {
		return true
	}
	return s == p
}
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				f[i][j] = i > 0 && match(s[i-1], p[j-1]) && f[i-1][j-1]
				continue
			}
			f[i][j] = f[i][j-1] || (i > 0 && f[i-1][j])
		}
	}
	return f[m][n]
}

// @lc code=end
func main() {
	fmt.Println(isMatch("aab", "a*?b"))
}
