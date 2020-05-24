package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == reverseInt(x) {
		return true
	}

	return false
}

func reverseInt(x int) int {
	flag := 1
	if x < 0 {
		flag, x = -1, -x
	}
	y := 0
	for x > 0 {
		y, x = y*10+x%10, x/10
	}

	y = flag * y
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}

	return y

}

// @lc code=end
