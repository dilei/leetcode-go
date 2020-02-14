package leetcode_go

import (
	"testing"
)

func Test_maxProduct(t *testing.T) {
	res := maxProduct([]int{2, 3, -2, 4})
	if res != 6 {
		t.Error(res)
	}
}
