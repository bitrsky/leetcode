package main

import "fmt"

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (40.75%)
 * Likes:    6560
 * Dislikes: 807
 * Total Accepted:    691.2K
 * Total Submissions: 1.7M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */

// @lc code=start

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) <= 0 {
		return ans
	}
	m, n := len(matrix), len(matrix[0])
	spiral := func(i1, j1, i2, j2 int) {

		for j := j1; j <= j2; j++ {
			ans = append(ans, matrix[i1][j])
		}
		for i := i1 + 1; i <= i2; i++ {
			ans = append(ans, matrix[i][j2])
		}
		if i1 != i2 {
			for j := j2 - 1; j >= j1; j-- {
				ans = append(ans, matrix[i2][j])
			}
		}
		if j1 != j2 {
			for i := i2 - 1; i > i1; i-- {
				ans = append(ans, matrix[i][j1])
			}
		}
	}

	for i1, j1, i2, j2 := 0, 0, m-1, n-1; i1 <= i2 && j1 <= j2; i1, j1, i2, j2 = i1+1, j1+1, i2-1, j2-1 {
		fmt.Println(i1, j1, i2, j2)
		spiral(i1, j1, i2, j2)
	}
	return ans
}

// @lc code=end
func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	// matrix := [][]int{{1}}
	ans := spiralOrder(matrix)
	fmt.Println(ans)
}
