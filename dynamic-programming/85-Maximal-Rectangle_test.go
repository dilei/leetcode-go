package dynamic_programming

import "testing"

func Test_maximalRectangle(t *testing.T) {
	nums := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	nums = [][]byte{
		{'1', '0', '1', '1', '1', '0', '0', '0', '1', '0'},
		{'0', '1', '0', '0', '0', '0', '0', '1', '1', '0'},
		{'0', '1', '0', '1', '0', '0', '0', '0', '1', '1'},
		{'1', '1', '1', '0', '0', '0', '0', '0', '1', '0'},
		{'0', '1', '1', '1', '0', '0', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '1', '0', '1', '1', '1', '0'},
	}
	res := maximalRectangle(nums)
	if res != 6 {
		t.Error(res)
	}
}
