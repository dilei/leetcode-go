package leetcodego

import (
	"fmt"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
