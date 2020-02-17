package backtracking

var keys = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	help("", digits, 0, &res)
	return res
}

func help(prefix string, digits string, offset int, res *[]string) {
	if offset >= len(digits) {
		*res = append(*res, prefix)
		return
	}

	str := keys[digits[offset]-'0']
	for i := 0; i < len(str); i++ {
		help(prefix+string(str[i]), digits, offset+1, res)
	}
}

// 第二种解法
func letterCombinations2(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	res = append(res, "")
	for i := 0; i < len(digits); i++ {
		x := digits[i] - '0'
		for len(res[0]) == i {
			prefix := res[0]
			res = res[1:]
			for _, ch := range keys[x] {
				res = append(res, prefix+string(ch))
			}
		}
	}
	return res
}
