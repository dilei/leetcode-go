package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_numDistinct(t *testing.T) {
	res := numDistinct("rabbbit", "rabbit")
	res2 := numDistinct("babgbag", "bag")
	fmt.Println(res, res2)
}
