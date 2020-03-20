package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	nums := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	res := minimumTotal(nums)
	fmt.Println(res)
}
