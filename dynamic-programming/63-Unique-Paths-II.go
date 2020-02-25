package dynamic_programming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// dp[i][j]表示到当前点的路径总数
	if obstacleGrid[0][0] == 1 {
		return 0
	} else {
		dp[0][0] = 1
	}

	// 第一行大于1的位置全是0
	flagPos := n // 保证不出现1时，下方循环赋值成功
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			flagPos = i
			break
		}
	}
	for i := 1; i < n; i++ {
		if i >= flagPos {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}

	flagPos = m
	// 第一列大于1的位置全是0
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			flagPos = i
			break
		}
	}
	for i := 1; i < m; i++ {
		if i >= flagPos {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 一维数组
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	w := len(obstacleGrid[0])
	dp := make([]int, w)
	dp[0] = 1
	for _, row := range obstacleGrid {
		for i := 0; i < w; i++ {
			if row[i] == 1 {
				dp[i] = 0
			} else if i > 0 {
				dp[i] += dp[i-1]
			}
		}
	}
	return dp[w-1]
}
