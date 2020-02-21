package backtracking

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}

	if p[1] != '*' {
		if len(s) == 0 {
			return false
		}
		return (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	// 处理相同前缀且含有*的情况 exp: s: aa p: a*
	for len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
		if isMatch(s, p[2:]) {
			return true
		}
		s = s[1:]
	}

	return isMatch(s, p[2:])
}

// 动态规划
func isMatch3(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	//    	 c * a * b
	//     0 1 2 3 4 5
	// 	 0 y n y n y n
	// a 1 n n n y y n
	// a 2 n n n n y n
	// b 3 n n n n n y
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	// 空字符匹配
	dp[0][0] = true
	for j := 2; j <= len(p); j++ {
		dp[0][j] = p[j-1] == '*' && dp[0][j-2]
	}

	for j := 1; j <= len(p); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || ((s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			}
		}
	}
	return dp[len(s)][len(p)]
}
