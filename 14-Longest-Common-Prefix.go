package leetcode_go

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], pre) != 0 {
			pre = pre[:len(pre)-1]
			if pre == "" {
				return ""
			}
		}
	}
	return pre
}
