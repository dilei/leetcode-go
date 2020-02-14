package leetcode_go

import "testing"

func Test_longestPalindrome(t *testing.T) {
	// res := longestPalindrome("abab")
	// if res != "aba" && res != "bab" {
	// 	t.Error(res)
	// }
	res := longestPalindrome("cbbd")
	if res != "bb" {
		t.Error(res)
	}
}
