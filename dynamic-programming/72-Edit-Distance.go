package dynamic_programming

import "math"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化第一行
	for i := 0; i < n; i++ {
		if word1[0] == word2[i] {
			dp[0][i] = i
		} else if i > 0 {
			dp[0][i] = dp[0][i-1] + 1
		} else {
			dp[0][i] = 1
		}
	}

	// 初始化第一列
	for i := 0; i < m; i++ {
		if word2[0] == word1[i] {
			dp[i][0] = i
		} else if i > 0 {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = 1
		}
	}

	// 填表
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m-1][n-1]
}

func min(x, y, z int) int {
	min := math.MaxInt32
	if x < min {
		min = x
	}
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

func minDistance2(word1 string, word2 string) int {
	m := len(word1) + 1
	n := len(word2) + 1
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 第一行
	for i := 0; i < n; i++ {
		dp[0][i] = i
	}

	// 第一列
	for j := 0; j < m; j++ {
		dp[j][0] = j
	}

	// 填表
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] { // 这里为什么是i-1和j-1？是为了让dp的索引和字符串的索引保持一致
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m-1][n-1]
}
