package leetcode_go

import "testing"

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	expect := 2
	res := majorityElement(nums)
	if res != expect {
		t.Error("unexpected")
	}
}
