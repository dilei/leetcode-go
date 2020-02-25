package dynamic_programming

func minPathSum(grid [][]int) int {
	w := len(grid[0])
	dp := make([]int, w)
	dp[0] = grid[0][0]
	for j, row := range grid {
		for i := 1; i < w; i++ {
			if j == 0 {
				dp[i] = row[i] + dp[i-1]
			} else if j > 0 {
				if i == 1 {
					dp[0] = dp[0] + row[0]
				}

				if dp[i-1] < dp[i] {
					dp[i] = row[i] + dp[i-1]
				} else {
					dp[i] += row[i]
				}
			}
		}
	}
	return dp[w-1]
}
