package backtracking

import (
	"fmt"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	res := permuteUnique([]int{1, 1, 2})
	if len(res) != 3 {
		t.Error(res)
	}
	fmt.Println(res)
}
