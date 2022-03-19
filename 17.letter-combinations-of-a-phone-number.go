package main

import "fmt"

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (52.91%)
 * Likes:    8900
 * Dislikes: 628
 * Total Accepted:    1.1M
 * Total Submissions: 2M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent. Return the answer in
 * any order.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * Example 2:
 *
 *
 * Input: digits = ""
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] is a digit in the range ['2', '9'].
 *
 *
 */

// @lc code=start
var num2char = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	ans := []string{}
	if len(digits) == 0 {
		return []string{}
	}
	str := num2char[digits[0]-'0']
	suffixs := letterCombinations(digits[1:])
	if len(suffixs) == 0 {
		suffixs = append(suffixs, "")
	}
	for _, s := range suffixs {
		for i := 0; i < len(str); i++ {
			ans = append(ans, string(append(append([]byte{}, str[i]), []byte(s)...)))
		}
	}
	return ans
}

// @lc code=end
func main() {

	fmt.Println(letterCombinations("23"))
}
