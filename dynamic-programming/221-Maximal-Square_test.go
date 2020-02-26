package dynamic_programming

import "testing"

func Test_maximalSquare(t *testing.T) {
	nums := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	res := maximalSquare(nums)
	if res != 4 {
		t.Error(res)
	}
}
