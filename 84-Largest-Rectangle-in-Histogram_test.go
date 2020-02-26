package leetcode_go

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	nums := []int{2, 1, 5, 6, 2, 3}
	res := largestRectangleArea(nums)
	if res != 10 {
		t.Error(res)
	}
}
