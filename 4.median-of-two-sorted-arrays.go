package main

import "fmt"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (33.60%)
 * Likes:    14861
 * Dislikes: 1872
 * Total Accepted:    1.3M
 * Total Submissions: 3.7M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return
 * the median of the two sorted arrays.
 *
 * The overall run time complexity should be O(log (m+n)).
 *
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,3], nums2 = [2]
 * Output: 2.00000
 * Explanation: merged array = [1,2,3] and median is 2.
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [1,2], nums2 = [3,4]
 * Output: 2.50000
 * Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 *
 *
 *
 * Constraints:
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lc code=start
func min(m, n int) int {
	if m > n {
		return n
	}
	return m
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 0 {
		return 0
	}
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	l, r, m1, m2 := 0, len(nums1), 0, 0

	for l <= r {
		m1 = (l + r) / 2
		m2 = (len(nums1)+len(nums2)+1)/2 - m1
		if m2 >= 1 && m1 < len(nums1) && nums2[m2-1] > nums1[m1] {
			l = m1 + 1
			continue
		}
		if m2 >= 0 && m1 >= 1 && nums1[m1-1] > nums2[m2] {
			r = m1 - 1
			continue
		}
		break
	}

	lmax, rmin := 0, 0
	if m1 == 0 {
		lmax = nums2[m2-1]
	} else if m2 == 0 {
		lmax = nums1[m1-1]
	} else {
		lmax = max(nums1[m1-1], nums2[m2-1])
	}
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(lmax * 1.0)
	}

	if m1 == len(nums1) {
		rmin = nums2[m2]
	} else if m2 == len(nums2) {
		rmin = nums1[m1]
	} else {
		rmin = min(nums1[m1], nums2[m2])
	}

	return float64(lmax+rmin) / 2
}

// @lc code=end
func main() {
	nums1 := []int{}
	nums2 := []int{1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
