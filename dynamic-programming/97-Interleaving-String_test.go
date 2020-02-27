package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_isInterleave(t *testing.T) {
	res := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	res2 := isInterleave("aabcc", "dbbca", "aadbbbaccc")
	fmt.Println(res, res2)
}
