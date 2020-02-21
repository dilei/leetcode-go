package backtracking

import (
	"fmt"
	"testing"
)

func Test_getPermutation(t *testing.T) {
	res := getPermutation2(3, 3)
	if res != "213" {
		t.Error(res)
	}
	fmt.Println(res)
}
