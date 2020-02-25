package dynamic_programming

import "testing"

func Test_climbStairs(t *testing.T) {
	res := climbStairs(3)
	if res != 3 {
		t.Error(res)
	}
}
