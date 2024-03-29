package main

/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (52.82%)
 * Likes:    4557
 * Dislikes: 144
 * Total Accepted:    323.1K
 * Total Submissions: 611.3K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * A sudoku solution must satisfy all of the following rules:
 *
 *
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 *
 *
 * The '.' character indicates empty cells.
 *
 *
 * Example 1:
 *
 *
 * Input: board =
 * [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
 * Output:
 * [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
 * Explanation: The input board is shown above and the only valid solution is
 * shown below:
 *
 *
 *
 *
 *
 * Constraints:
 *
 *
 * board.length == 9
 * board[i].length == 9
 * board[i][j] is a digit or '.'.
 * It is guaranteed that the input board has only one solution.
 *
 *
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	var row, col [9]int
	var box [3][3]int
	var space [][2]int

	filp := func(i, j int, digit byte) {
		row[i] ^= 1 << digit
		col[j] ^= 1 << digit
		box[i/3][j/3] ^= 1 << digit
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				space = append(space, [2]int{i, j})
				continue
			}
			digit := board[i][j] - '1'
			filp(i, j, digit)
		}
	}

	var backtrack func(int) bool

	backtrack = func(pos int) bool {
		if pos == len(space) {
			return true
		}
		i, j := space[pos][0], space[pos][1]
		mask := 0x1ff ^ uint(row[i]|col[j]|box[i/3][j/3])
		for digit := byte(0); mask > 0; digit, mask = digit+1, mask/2 {
			if mask%2 == 0 {
				continue
			}
			// digit := byte(bits.TrailingZeros(mask))
			filp(i, j, digit)
			board[i][j] = digit + '1'
			if backtrack(pos + 1) {
				return true
			}
			board[i][j] = '.'
			filp(i, j, digit)
		}
		return false
	}
	backtrack(0)

}

// @lc code=end
func main() {

}
