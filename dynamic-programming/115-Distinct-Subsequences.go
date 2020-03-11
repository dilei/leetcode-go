package dynamic_programming

func numDistinct(s string, t string) int {
	sLen := len(s)
	tLen := len(t)

	dp := make([][]int, tLen+1)
	for i := range dp {
		dp[i] = make([]int, sLen+1)
	}
	dp[0][0] = 1

	// t =="" 时，每个s[i]都有一种匹配
	for i := 1; i < len(s); i++ {
		dp[0][i] = 1
	}

	// s==""时，每个t[i]都不能匹配, default 0

	for i := 1; i <= tLen; i++ {
		for j := 1; j <= sLen; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1] // 认为是前匹配的数字
			}
		}
	}
	return dp[tLen][sLen]
}

// 超时
func numDistinct2(s string, t string) int {
	if len(s) < len(t) || len(s) == 0 {
		return 0
	}

	res := 0
	help115(s, t, &res)
	return res
}

func help115(s string, t string, res *int) {
	if len(t) == 0 {
		*res++
		return
	}
	if len(s) < len(t) {
		return
	}
	if s[0] == t[0] {
		help115(s[1:], t[1:], res)
	}
	help115(s[1:], t, res)
}
