package divide_conquer

import (
	"fmt"
	"testing"
)

func Test_addOperators(t *testing.T) {
	res := addOperators("123", 6)
	res2 := addOperators("232", 8)
	fmt.Println(res, res2)
}
