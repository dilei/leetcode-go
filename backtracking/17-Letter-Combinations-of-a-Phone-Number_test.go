package backtracking

import (
	"fmt"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	// res := letterCombinations("23")
	res := letterCombinations2("23")
	if len(res) != 9 {
		t.Error("error")
	}
	fmt.Println(res)
}
