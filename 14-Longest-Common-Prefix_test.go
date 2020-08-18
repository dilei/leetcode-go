package leetcodego

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
