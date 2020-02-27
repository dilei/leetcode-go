package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_numDecodings(t *testing.T) {
	res := numDecodings("12")
	res2 := numDecodings("226")
	fmt.Println(res, res2)
}
