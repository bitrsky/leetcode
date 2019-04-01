package leetcode
/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (28.85%)
 * Total Accepted:    99.4K
 * Total Submissions: 344.6K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 
 * 示例 1:
 * 
 * 输入: "abcabcbb"
 * 输出: 3 
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 
 * 
 * 示例 2:
 * 
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 
 * 
 * 示例 3:
 * 
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 
 * 
 */
 //lastLocation[v] 表示字符v上一次出现的位置
 // left 不重复的第一个字符的位置
func lengthOfLongestSubstring(s string) int {
	lastLocation,max,left := map[rune]int{},0,0
	for index,v := range s {
		if b,ok := lastLocation[v]; ok && b >= left {
			if index - left > max {
				max = index - left
			}
			left = b+1
		}
		lastLocation[v] = index
	}
	if len(s) - left > max {
		max = len(s) - left
	}
	return max
}

