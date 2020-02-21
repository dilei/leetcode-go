package dynamic_programming

import "testing"

func Test_longestPalindrome(t *testing.T) {
	// res := longestPalindrome("abab")
	// if res != "aba" && res != "bab" {
	// 	t.Error(res)
	// }
	res := longestPalindrome2("cbbd")
	if res != "bb" {
		t.Error(res)
	}
}
