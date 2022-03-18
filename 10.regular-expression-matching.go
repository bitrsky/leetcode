package main

import "fmt"

/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (28.17%)
 * Likes:    7730
 * Dislikes: 1129
 * Total Accepted:    656.2K
 * Total Submissions: 2.3M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string s and a pattern p, implement regular expression
 * matching with support for '.' and '*' where:
 *
 *
 * '.' Matches any single character.​​​​
 * '*' Matches zero or more of the preceding element.
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
 * Input: s = "aa", p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the preceding element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 *
 *
 * Example 3:
 *
 *
 * Input: s = "ab", p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 20
 * 1 <= p.length <= 30
 * s contains only lowercase English letters.
 * p contains only lowercase English letters, '.', and '*'.
 * It is guaranteed for each appearance of the character '*', there will be a
 * previous valid character to match.
 *
 *
 */

// @lc code=start

func match(s, p byte) bool {
	if p == '.' {
		return true
	}
	return s == p
}

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}

	m := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		m[i] = make([]bool, len(p)+1)
	}
	m[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] != '*' {
				m[i][j] = i > 0 && match(s[i-1], p[j-1]) && m[i-1][j-1]
				continue
			}
			m[i][j] = m[i][j] || m[i][j-2]
			if i > 0 && match(s[i-1], p[j-2]) {
				m[i][j] = m[i][j] || m[i-1][j]
			}
		}
	}
	return m[len(s)][len(p)]
}

// @lc code=end
func main() {

	fmt.Println(isMatch("11", "1*"))

}
