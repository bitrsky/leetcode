//package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (32.80%)
 * Likes:    21178
 * Dislikes: 950
 * Total Accepted:    3M
 * Total Submissions: 9M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string s, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * Notice that the answer must be a substring, "pwke" is a subsequence and not
 * a substring.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s consists of English letters, digits, symbols and spaces.
 *
 *
 */

/*
* Node: record the position of the first character of a substring without repeating character
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var loc [128]int
	max, start := 0, 0
	for i := 1; i <= len(s); i++ {
		ch := int(s[i-1])
		if start < loc[ch] {
			start = loc[ch]
		}

		if max < i-start {
			max = i - start
		}

		loc[ch] = i
	}
	return max
}

/*
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
*/

// @lc code=end
