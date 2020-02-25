package dynamic_programming

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	nums := []int{10, 15, 20}
	res := minCostClimbingStairs(nums)
	if res != 15 {
		t.Error(res)
	}
}
