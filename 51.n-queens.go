package main

import "fmt"

/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (56.14%)
 * Likes:    5146
 * Dislikes: 144
 * Total Accepted:    340.4K
 * Total Submissions: 605.1K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n x n
 * chessboard such that no two queens attack each other.
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 * You may return the answer in any order.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space,
 * respectively.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4
 * Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: [["Q"]]
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
func generateMap(n int, queens [][2]int) []string {
	m := make([][]byte, 0, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		m = append(m, row)
	}
	for _, loc := range queens {
		m[loc[0]][loc[1]] = 'Q'
	}
	ans := []string{}
	for i := 0; i < n; i++ {
		ans = append(ans, string(m[i]))
	}
	return ans
}
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
func solveNQueens(n int) [][]string {
	var queens [][2]int
	ans := [][]string{}
	var backtrack func(int)

	backtrack = func(row int) {
		if row == n {
			ans = append(ans, generateMap(n, queens))
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

	ans := solveNQueens(4)
	fmt.Println(ans)
}
