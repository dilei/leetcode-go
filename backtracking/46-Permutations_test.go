package backtracking

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	res := permute([]int{0, 2, 3})
	if len(res) != 5 {
		t.Error(res)
	}
	fmt.Println(res)
}
