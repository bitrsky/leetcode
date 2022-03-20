package main

import "fmt"

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (62.04%)
 * Likes:    2635
 * Dislikes: 163
 * Total Accepted:    309.4K
 * Total Submissions: 498.6K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate an n x n matrix filled with elements
 * from 1 to n^2 in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3
 * Output: [[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: [[1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 20
 *
 *
 */

// @lc code=start
func generate(ans [][]int, i1, j1, i2, j2, num int) int {
	for j := j1; j <= j2; j++ {
		ans[i1][j] = num
		num++
	}
	for i := i1 + 1; i <= i2; i++ {
		ans[i][j2] = num
		num++
	}
	for j := j2 - 1; j >= j1; j-- {
		ans[i2][j] = num
		num++
	}
	for i := i2 - 1; i > i1; i-- {
		ans[i][j1] = num
		num++
	}
	return num
}
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	num := 1
	for i1, j1, i2, j2 := 0, 0, n-1, n-1; i1 <= i2 && j1 <= j2; i1, j1, i2, j2 = i1+1, j1+1, i2-1, j2-1 {
		num = generate(ans, i1, j1, i2, j2, num)
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(generateMatrix(3))
}
