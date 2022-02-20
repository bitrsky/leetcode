package main

import "fmt"

/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 *
 * https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (27.72%)
 * Likes:    1745
 * Dislikes: 1762
 * Total Accepted:    240.8K
 * Total Submissions: 868.6K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * You are given a string s and an array of strings words of the same length.
 * Return all starting indices of substring(s) in s that is a concatenation of
 * each word in words exactly once, in any order, and without any intervening
 * characters.
 *
 * You can return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "barfoothefoobarman", words = ["foo","bar"]
 * Output: [0,9]
 * Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar"
 * respectively.
 * The output order does not matter, returning [9,0] is fine too.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 * Output: [6,9,12]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of lower-case English letters.
 * 1 <= words.length <= 5000
 * 1 <= words[i].length <= 30
 * words[i] consists of lower-case English letters.
 *
 *
 */

// @lc code=start
func isMatch(s string, wordsMap map[string]int, wlen int) bool {
	contain := make(map[string]int)
	for i := 0; i < len(s); i += wlen {
		word := s[i : i+wlen]
		limit, ok := wordsMap[word]
		if !ok {
			return false
		}
		k, ok := contain[word]
		if ok {
			contain[word] = k + 1
		} else {
			contain[word] = 1
		}
		if contain[word] > limit {
			return false
		}
	}
	return true
}

func findSubstring(s string, words []string) []int {
	ans := []int{}
	wordsMap := make(map[string]int)
	wordLen := len(words[0])
	for i := 0; i < len(words); i += 1 {
		if c, ok := wordsMap[words[i]]; ok {
			wordsMap[words[i]] = c + 1
			continue
		}
		wordsMap[words[i]] = 1
	}
	wordsLen := wordLen * len(words)
	for i := 0; i <= len(s)-wordsLen; i += 1 {
		if _, ok := wordsMap[s[i:i+wordLen]]; !ok {
			continue
		}
		if isMatch(s[i:i+wordsLen], wordsMap, wordLen) {
			ans = append(ans, i)
		}
	}
	return ans
}

// @lc code=end

// Input: s = "barfoothefoobarman", words = ["foo","bar"]
// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

func main() {
	// s := "wordgoodgoodgoodbestword"
	// words := []string{"word", "good", "best", "good"}
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
}
