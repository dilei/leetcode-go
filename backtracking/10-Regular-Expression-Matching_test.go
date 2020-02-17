package backtracking

import "testing"

func Test_isMatch(t *testing.T) {
	// res := isMatch("aab", "c*a*b")
	res := isMatch("aa", "a*")
	if res != true {
		t.Error("error")
	}
}
