package dynamic_programming

func isInterleave(s1 string, s2 string, s3 string) bool {
	ls1 := len(s1)
	ls2 := len(s2)
	ls3 := len(s3)
	if ls1+ls2 != ls3 {
		return false
	}
	if ls1+ls2 == 0 {
		return true
	}
	// dp[i][j]:表示s1[0...i]和s2[0...j]和s3[0...i+j-1]匹配
	dp := make([][]bool, ls1+1)
	for i := range dp {
		dp[i] = make([]bool, ls2+1)
	}
	dp[0][0] = true

	// 初始化第一行
	for i := 1; i <= ls2; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}
	// 初始化第一列
	for i := 1; i <= ls1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	// 填表
	for i := 1; i <= ls1; i++ {
		for j := 1; j <= ls2; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[ls1][ls2]
}

// 2
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1)+len(s2) == 0 {
		return true
	}
	return help97(s1, s2, s3)
}

func help97(s1 string, s2 string, s3 string) bool {
	if len(s3) == 0 {
		return true
	}
	if s1 != "" && s3[0] == s1[0] {
		if help97(s1[1:], s2, s3[1:]) {
			return true
		}
	}
	if s2 != "" && s3[0] == s2[0] {
		if help97(s1, s2[1:], s3[1:]) {
			return true
		}
	}
	return false
}
