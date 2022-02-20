package main

import "fmt"

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (69.30%)
 * Likes:    11597
 * Dislikes: 452
 * Total Accepted:    973.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '3'
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 * Example 1:
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 * Input: n = 1
 * Output: ["()"]
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 8
 *
 *
 */

/*
 * Node : left <= right
 */
// @lc code=start
func generate(cur string, left, right int, ans *[]string) {
	if left == 0 && right == 0 {
		*ans = append(*ans, cur)
		return
	}
	if left > 0 {
		generate(cur+"(", left-1, right, ans)
	}
	if left < right {
		generate(cur+")", left, right-1, ans)
	}
}

func generateParenthesis(n int) []string {
	ans := []string{}
	generate("(", n-1, n, &ans)
	return ans
}

// @lc code=end

func main() {
	fmt.Println(generateParenthesis(5))
}
