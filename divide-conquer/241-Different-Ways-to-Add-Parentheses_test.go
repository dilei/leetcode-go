package divide_conquer

import (
	"fmt"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	res := diffWaysToCompute("2-1-1")
	res2 := diffWaysToCompute("2*3-4*5")
	fmt.Println(res, res2)
}
