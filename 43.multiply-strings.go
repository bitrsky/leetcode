package main

import "fmt"

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (37.29%)
 * Likes:    3993
 * Dislikes: 1589
 * Total Accepted:    479K
 * Total Submissions: 1.3M
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 *
 * Note: You must not use any built-in BigInteger library or convert the inputs
 * to integer directly.
 *
 *
 * Example 1:
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 * Example 2:
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 *
 * Constraints:
 *
 *
 * 1 <= num1.length, num2.length <= 200
 * num1 and num2 consist of digits only.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 *
 *
 */

// @lc code=start

func s2b(num string) []uint8 {
	n := make([]uint8, 0, len(num))
	for i := len(num) - 1; i >= 0; i-- {
		n = append(n, num[i]-'0')
	}
	return n
}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := s2b(num1), s2b(num2)
	num := make([]uint8, len(n1)+len(n2), len(n1)+len(n2))
	for i := 0; i < len(n1); i++ {
		var carry uint8
		j := 0
		for ; j < len(n2); j++ {
			num[i+j] += n1[i]*n2[j] + carry
			carry = num[i+j] / 10
			num[i+j] = num[i+j] % 10
		}
		for ; carry > 0; j++ {
			num[i+j] += carry
			carry = num[i+j] / 10
			num[i+j] = num[i+j] % 10
		}
	}
	ans := []byte{}
	i := len(num) - 1
	for ; i >= 0 && num[i] == 0; i-- {
	}
	for ; i >= 0; i-- {
		ans = append(ans, '0'+num[i])
	}
	return string(ans)
}

// @lc code=end
func main() {
	num1 := "123"
	num2 := "369"
	fmt.Println(multiply(num1, num2))
}
