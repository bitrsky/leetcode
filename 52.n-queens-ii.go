package main

import "fmt"

/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 *
 * https://leetcode.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (65.66%)
 * Likes:    1609
 * Dislikes: 209
 * Total Accepted:    212.5K
 * Total Submissions: 323.2K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n x n
 * chessboard such that no two queens attack each other.
 *
 * Given an integer n, return the number of distinct solutions to the n-queens
 * puzzle.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4
 * Output: 2
 * Explanation: There are two distinct solutions to the 4-queens puzzle as
 * shown.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 9
 *
 *
 */

// @lc code=start

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func isAttack(queens [][2]int, i, j int) bool {
	for _, loc := range queens {
		if j == loc[1] {
			return true
		}
		if abs(i-loc[0]) == abs(j-loc[1]) {
			return true
		}
	}
	return false
}

func totalNQueens(n int) int {
	var queens [][2]int
	ans := 0
	var backtrack func(int)

	backtrack = func(row int) {
		if row == n {
			ans += 1
			return
		}
		for i := 0; i < n; i++ {
			if isAttack(queens, row, i) {
				continue
			}
			queens = append(queens, [2]int{row, i})
			backtrack(row + 1)
			queens = queens[:len(queens)-1]
		}
	}

	backtrack(0)
	return ans

}

// @lc code=end
func main() {
	fmt.Println(totalNQueens(4))
}
