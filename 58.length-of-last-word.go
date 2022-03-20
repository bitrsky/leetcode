package main

import "fmt"

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (36.85%)
 * Likes:    777
 * Dislikes: 59
 * Total Accepted:    695.8K
 * Total Submissions: 1.9M
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consisting of some words separated by some number of
 * spaces, return the length of the last word in the string.
 *
 * A word is a maximal substring consisting of non-space characters only.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "Hello World"
 * Output: 5
 * Explanation: The last word is "World" with length 5.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "   fly me   to   the moon  "
 * Output: 4
 * Explanation: The last word is "moon" with length 4.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "luffy is still joyboy"
 * Output: 6
 * Explanation: The last word is "joyboy" with length 6.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of only English letters and spaces ' '.
 * There will be at least one word in s.
 *
 *
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	end := i
	for ; i >= 0 && s[i] != ' '; i-- {
	}
	return end - i
}

// @lc code=end
func main() {
	fmt.Println(lengthOfLastWord("qwe ewrw rerd"))
}
