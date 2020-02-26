package dynamic_programming

import "testing"

func Test_maxProfit(t *testing.T) {
	res := maxProfit2([]int{7, 1, 5, 3, 6, 4})
	if res != 5 {
		t.Error(res)
	}
}
