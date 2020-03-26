package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_maxProfit123(t *testing.T) {
	nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	res := maxProfit123(nums)
	fmt.Println(res)
}
