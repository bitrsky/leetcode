package main

import "fmt"

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (40.87%)
 * Likes:    3322
 * Dislikes: 7737
 * Total Accepted:    710.3K
 * Total Submissions: 1.7M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * Example 3:
 *
 *
 * Input: s = "A", numRows = 1
 * Output: "A"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consists of English letters (lower-case and upper-case), ',' and '.'.
 * 1 <= numRows <= 1000
 *
 *
 */

// @lc code=start
func convert(s string, numRows int) string {
	ls := len(s)
	if numRows == 1 || ls <= numRows {
		return s
	}

	ans := make([]byte, 0, ls)
	mod := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < ls; j += mod {
			ans = append(ans, s[j+i])
			if i > 0 && i < numRows-1 && j+mod-i < ls {
				ans = append(ans, s[j+mod-i])
			}
		}
	}
	return string(ans)
}

// @lc code=end
func main() {
	fmt.Println(convert("QA", 1))

}
