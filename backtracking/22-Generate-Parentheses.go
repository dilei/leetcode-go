package backtracking

func generateParenthesis(n int) []string {
	var res []string
	help22(&res, "", 0, 0, n)
	return res
}

func help22(res *[]string, str string, open, close, max int) {
	if len(str) == max*2 {
		*res = append(*res, str)
		return
	}

	if open < max { // 还能放入“（”
		help22(res, str+"(", open+1, close, max)
	}
	if close < open { // 还能放入“）”
		help22(res, str+")", open, close+1, max)
	}
}
