package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	res := generateTrees(3)
	for _, val := range res {
		fmt.Println(val)
	}
}
