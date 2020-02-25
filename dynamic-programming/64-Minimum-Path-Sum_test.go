package dynamic_programming

import "testing"

func Test_minPathSum(t *testing.T) {
	nums := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := minPathSum(nums)
	if res != 7 {
		t.Error(res)
	}
}
