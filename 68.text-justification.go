package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 *
 * https://leetcode.com/problems/text-justification/description/
 *
 * algorithms
 * Hard (34.65%)
 * Likes:    1669
 * Dislikes: 2774
 * Total Accepted:    235.5K
 * Total Submissions: 679.2K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * Given an array of strings words and a width maxWidth, format the text such
 * that each line has exactly maxWidth characters and is fully (left and right)
 * justified.
 *
 * You should pack your words in a greedy approach; that is, pack as many words
 * as you can in each line. Pad extra spaces ' ' when necessary so that each
 * line has exactly maxWidth characters.
 *
 * Extra spaces between words should be distributed as evenly as possible. If
 * the number of spaces on a line does not divide evenly between words, the
 * empty slots on the left will be assigned more spaces than the slots on the
 * right.
 *
 * For the last line of text, it should be left-justified and no extra space is
 * inserted between words.
 *
 * Note:
 *
 *
 * A word is defined as a character sequence consisting of non-space characters
 * only.
 * Each word's length is guaranteed to be greater than 0 and not exceed
 * maxWidth.
 * The input array words contains at least one word.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["This", "is", "an", "example", "of", "text",
 * "justification."], maxWidth = 16
 * Output:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 * Example 2:
 *
 *
 * Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth
 * = 16
 * Output:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * Explanation: Note that the last line is "shall be    " instead of "shall
 * be", because the last line must be left-justified instead of
 * fully-justified.
 * Note that the second line is also left-justified becase it contains only one
 * word.
 *
 * Example 3:
 *
 *
 * Input: words =
 * ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"],
 * maxWidth = 20
 * Output:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 * Constraints:
 *
 *
 * 1 <= words.length <= 300
 * 1 <= words[i].length <= 20
 * words[i] consists of only English letters and symbols.
 * 1 <= maxWidth <= 100
 * words[i].length <= maxWidth
 *
 *
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	ans := []string{}
	l, r, n := 0, 0, len(words)
	for {
		l = r
		lg := 0
		for ; r < n && lg+len(words[r])+r-l <= maxWidth; r++ {
			lg += len(words[r])
		}
		if r == n {
			s := strings.Join(words[l:], " ")
			ans = append(ans, s+strings.Repeat(" ", maxWidth-len(s)))
			break
		}
		wordsNum := r - l
		spaceNum := maxWidth - lg
		if wordsNum == 1 {
			ans = append(ans, words[l]+strings.Repeat(" ", spaceNum))
			continue
		}
		avgSpace := spaceNum / (wordsNum - 1)
		extraSpace := spaceNum % (wordsNum - 1)
		s1 := strings.Join(words[l:l+extraSpace+1], strings.Repeat(" ", avgSpace+1))
		s2 := strings.Join(words[l+extraSpace+1:r], strings.Repeat(" ", avgSpace))
		ans = append(ans, s1+strings.Repeat(" ", avgSpace)+s2)
	}
	return ans
}

// @lc code=end
func main() {
	// words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	ans := fullJustify(words, 16)
	fmt.Println(ans)
}
