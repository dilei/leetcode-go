package leetcode_go

import "testing"

func Test_romanToInt(t *testing.T) {
	res := romanToInt("DCXXI")
	if res != 621 {
		t.Error(res)
	}
}
