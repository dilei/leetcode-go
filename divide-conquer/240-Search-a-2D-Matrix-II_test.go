package divide_conquer

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	nums := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	res := searchMatrix(nums, 5)
	fmt.Println(res)
}
