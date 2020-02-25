package dynamic_programming

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1 // dp[i][j]表示到当前点的路径总数
	// 第一行每个位置只有一种路线
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}
	// 第一列每个位置只有一种路线
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 一维数组
func uniquePaths2(m, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[n-1]
}
