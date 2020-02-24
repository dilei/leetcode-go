package backtracking

func isMatch2(s string, p string) bool {
	var si, pi, match int
	var startIndex = -1
	for si < len(s) {
		if pi < len(p) && (p[pi] == s[si] || p[pi] == '?') {
			si++
			pi++
		} else if pi < len(p) && p[pi] == '*' {
			startIndex = pi
			match = si
			pi++
		} else if startIndex != -1 {
			pi = startIndex + 1
			match++
			si = match
		} else {
			return false
		}
	}

	for pi < len(p) && p[pi] == '*' {
		pi++
	}
	return pi == len(p)
}

// 动态规划
func isMatchDynamic(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	// dp[i][j]表示 s[0...i-1]和p[0...j-1]是否匹配
	dp[0][0] = true // s:"" p:""

	// 初始化dp[i][0],全为false
	// 默认值以满足fasle

	// 第一行dp[0][j], 只有 "*","**", "***"满足需求
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else { // 不是连续的*，就返回
			break
		}
	}

	// 填表
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
		}
	}
	return dp[len(s)][len(p)]
}
