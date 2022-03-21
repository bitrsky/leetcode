package main

import "fmt"

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (42.29%)
 * Likes:    6363
 * Dislikes: 248
 * Total Accepted:    683.8K
 * Total Submissions: 1.6M
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * Write an efficient algorithm that searches for a value target in an m x n
 * integer matrix matrix. This matrix has the following properties:
 *
 *
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 *
 *
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for row < m && col >= 0 {
		if matrix[row][n-1] < target {
			row++
			continue
		}
		if matrix[row][col] > target {
			col--
			continue
		}
		if matrix[row][col] == target {
			return true
		}
		break
	}
	return false
}

// @lc code=end
func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(searchMatrix(matrix, 6))
}
