package divide_conquer

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}
