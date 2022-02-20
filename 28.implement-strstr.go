package main

import "fmt"

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (35.66%)
 * Likes:    3530
 * Dislikes: 3313
 * Total Accepted:    1.1M
 * Total Submissions: 3.2M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 *
 * Example 1:
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 * Example 2:
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 * Example 3:
 * Input: haystack = "", needle = ""
 * Output: 0
 *
 *
 * Constraints:
 *
 *
 * 0 <= haystack.length, needle.length <= 5 * 10^4
 * haystack and needle consist of only lower-case English characters.
 *
 *
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" || needle == haystack {
		return 0
	}
	ans := -1
	if len(needle) > len(haystack) {
		return ans
	}
	for i := 0; i < len(haystack); i += 1 {
		if needle[0] != haystack[i] {
			continue
		}
		if len(needle) > len(haystack)-i {
			break
		}
		if haystack[i:i+len(needle)] == needle {
			ans = i
			break
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(strStr("", ""))
}
