package dynamic_programming

import (
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	// res := longestValidParentheses("(()")
	res := longestValidParentheses2("()")
	if res != 2 {
		t.Error(res)
	}
}
