package backtracking

import (
	"fmt"
	"testing"
)

func Test_combinationSumII(t *testing.T) {
	res := combinationSumII([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	if len(res) != 4 {
		t.Error(res)
	}
	fmt.Println(res)
}
