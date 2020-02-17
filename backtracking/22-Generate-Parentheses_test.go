package backtracking

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	res := generateParenthesis(3)
	if len(res) != 5 {
		t.Error(res)
	}
	fmt.Println(res)
}
