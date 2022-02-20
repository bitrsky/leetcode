package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (63.32%)
 * Likes:    8281
 * Dislikes: 286
 * Total Accepted:    1.2M
 * Total Submissions: 1.9M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings strs, group the anagrams together. You can return
 * the answer in any order.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * Input: strs = ["a"]
 * Output: [["a"]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lowercase English letters.
 *
 *
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		counter := make([]int, 26, 26)
		for i := 0; i < len(str); i++ {
			counter[str[i]-'a'] += 1
		}
		sb := []byte{}
		for i := 0; i < len(counter); i++ {
			sb = append(sb, byte('a'+i))
			sb = append(sb, []byte(strconv.Itoa(counter[i]))...)
		}
		st := string(sb)
		if _, ok := m[st]; !ok {
			m[st] = []string{}
		}
		m[st] = append(m[st], str)
	}
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end
func main() {
	strs := []string{"abc", "cdv"}
	fmt.Println(groupAnagrams(strs))
}
