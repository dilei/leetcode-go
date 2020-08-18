package leetcodego

import (
	"fmt"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	res := intToRoman(3)
	res1 := intToRoman(4)
	res2 := intToRoman(9)
	res3 := intToRoman(58)
	res4 := intToRoman(1994)
	fmt.Println(res, res1, res2, res3, res4)
}
