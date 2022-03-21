package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (39.34%)
 * Likes:    8766
 * Dislikes: 336
 * Total Accepted:    919.2K
 * Total Submissions: 2.3M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given an m x n grid of characters board and a string word, return true if
 * word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cells,
 * where adjacent cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once.
 *
 *
 * Example 1:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCCED"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "SEE"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCB"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n = board[i].length
 * 1 <= m, n <= 6
 * 1 <= word.length <= 15
 * board and word consists of only lowercase and uppercase English letters.
 *
 *
 *
 * Follow up: Could you use search pruning to make your solution faster with a
 * larger board?
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	var backtrack func(i, j int, seq []byte) bool
	backtrack = func(i, j int, seq []byte) bool {
		visit[i][j] = true
		seq = append(seq, board[i][j])
		defer func() {
			visit[i][j] = false
			seq = seq[:len(seq)-1]
		}()
		if string(seq) == word {
			return true
		}
		if !strings.HasPrefix(word, string(seq)) {
			return false
		}
		if i > 0 && !visit[i-1][j] {
			if backtrack(i-1, j, seq) {
				return true
			}
		}
		if i < m-1 && !visit[i+1][j] {
			if backtrack(i+1, j, seq) {
				return true
			}
		}

		if j > 0 && !visit[i][j-1] {
			if backtrack(i, j-1, seq) {
				return true
			}
		}

		if j < n-1 && !visit[i][j+1] {
			if backtrack(i, j+1, seq) {
				return true
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] != board[i][j] {
				continue
			}
			if backtrack(i, j, make([]byte, 0, len(word))) {
				return true
			}
		}
	}
	return false

}

// @lc code=end
func main() {
	board := [][]byte{{'1', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}}
	fmt.Println(exist(board, "129"))
}
