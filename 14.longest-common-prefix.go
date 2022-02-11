package main

import "fmt"

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (38.84%)
 * Likes:    6876
 * Dislikes: 2783
 * Total Accepted:    1.4M
 * Total Submissions: 3.6M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] consists of only lower-case English letters.
 *
 *
 */

// @lc code=start
func minStr(strs []string) string {
	minStr := strs[0]
	for _, str := range strs {
		if len(minStr) > len(str) {
			minStr = str
		}
	}
	return minStr
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minStr := minStr(strs)
	var index int
	for index = 0; index < len(minStr); index++ {
		isContain := true
		for _, str := range strs {
			if str[index] != minStr[index] {
				isContain = false
				break
			}
		}
		if !isContain {
			break
		}
	}
	return minStr[0:index]
}

// @lc code=end

func main() {
	fmt.Println(longestCommonPrefix([]string{"abcd"}))
}
