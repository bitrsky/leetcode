package leetcode

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9
	got := twoSum(nums, target)
	fmt.Println(got)
}
