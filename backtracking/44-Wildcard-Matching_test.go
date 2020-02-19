package backtracking

import "testing"

func Test_isMatch2(t *testing.T) {
	// res := isMatch2("adceb", "*a*b")
	res := isMatch2("adcebb", "*a*b")
	if res == false {
		t.Error(res)
	}
}
