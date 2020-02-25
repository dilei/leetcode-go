package dynamic_programming

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	// nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	nums := []int{1, 2}
	res := maxSubArray2(nums)
	if res != 3 {
		t.Error(res)
	}
}
