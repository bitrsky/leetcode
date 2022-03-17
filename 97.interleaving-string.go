package main

import "fmt"

/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 *
 * https://leetcode.com/problems/interleaving-string/description/
 *
 * algorithms
 * Medium (34.55%)
 * Likes:    3664
 * Dislikes: 196
 * Total Accepted:    254.4K
 * Total Submissions: 736K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * Given strings s1, s2, and s3, find whether s3 is formed by an interleaving
 * of s1 and s2.
 *
 * An interleaving of two strings s and t is a configuration where they are
 * divided into non-empty substrings such that:
 *
 *
 * s = s1 + s2 + ... + sn
 * t = t1 + t2 + ... + tm
 * |n - m| <= 1
 * The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 +
 * t3 + s3 + ...
 *
 *
 * Note: a + b is the concatenation of strings a and b.
 *
 *
 * Example 1:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: s1 = "", s2 = "", s3 = ""
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s1.length, s2.length <= 100
 * 0 <= s3.length <= 200
 * s1, s2, and s3 consist of lowercase English letters.
 *
 *
 *
 * Follow up: Could you solve it using only O(s2.length) additional memory
 * space?
 *
 */

// @lc code=start

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n+m != t {
		return false
	}
	isOk := make([]bool, m+1)
	isOk[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			v := i + j
			if i > 0 {
				isOk[j] = isOk[j] && s1[i-1] == s3[v-1]
			}
			if j > 0 {
				isOk[j] = isOk[j] || (isOk[j-1] && s2[j-1] == s3[v-1])
			}
		}
	}
	return isOk[m]
}
func isInterleave1(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n+m != t {
		return false
	}
	isOk := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		isOk[i] = make([]bool, m+1)
	}
	isOk[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			v := i + j
			if i > 0 {
				isOk[i][j] = isOk[i][j] || isOk[i-1][j] && s1[i-1] == s3[v-1]
			}
			if j > 0 {
				isOk[i][j] = isOk[i][j] || isOk[i][j-1] && s2[j-1] == s3[v-1]
			}
		}
	}
	return isOk[n][m]
}

// @lc code=end

func main() {
	s1 := "123"
	s2 := "123"
	s3 := "112233"
	fmt.Println(isInterleave(s1, s2, s3))
}
