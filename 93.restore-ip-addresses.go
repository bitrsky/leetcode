package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (41.34%)
 * Likes:    2579
 * Dislikes: 621
 * Total Accepted:    289K
 * Total Submissions: 697.9K
 * Testcase Example:  '"25525511135"'
 *
 * A valid IP address consists of exactly four integers separated by single
 * dots. Each integer is between 0 and 255 (inclusive) and cannot have leading
 * zeros.
 *
 *
 * For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but
 * "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP
 * addresses.
 *
 *
 * Given a string s containing only digits, return all possible valid IP
 * addresses that can be formed by inserting dots into s. You are not allowed
 * to reorder or remove any digits in s. You may return the valid IP addresses
 * in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "25525511135"
 * Output: ["255.255.11.135","255.255.111.35"]
 *
 *
 * Example 2:
 *
 *
 * Input: s = "0000"
 * Output: ["0.0.0.0"]
 *
 *
 * Example 3:
 *
 *
 * Input: s = "101023"
 * Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 20
 * s consists of digits only.
 *
 *
 */

// @lc code=start
func validNum(s string) bool {
	if len(s) == 1 {
		return true
	}
	if s[0] == '0' {
		return false
	}
	if v, _ := strconv.Atoi(s); v <= 255 {
		return true
	}
	return false
}
func divide(s string, points int) []string {
	if points == 0 {
		if validNum(s) {
			return []string{s}
		}
		return []string{}
	}
	ans := []string{}
	if s == "" {
		return ans
	}
	for i := 0; i < 3 && i < len(s)-1; i++ {
		if !validNum(s[:i+1]) {
			continue
		}
		for _, ls := range divide(s[i+1:], points-1) {
			ans = append(ans, fmt.Sprintf("%s.%s", s[:i+1], ls))
		}
	}
	return ans
}
func restoreIpAddresses(s string) []string {
	return divide(s, 3)
}

// @lc code=end
