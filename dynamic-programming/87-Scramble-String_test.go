package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_isScramble(t *testing.T) {
	res := isScramble("great", "rgeat")
	res2 := isScramble("abcde", "caebd")
	fmt.Println(res, res2)
}
