package backtracking

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	res := combinationSum([]int{2, 3, 5}, 8)
	if len(res) != 3 {
		t.Error(res)
	}
	fmt.Println(res)
}
