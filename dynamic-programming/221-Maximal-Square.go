package dynamic_programming

import "math"

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	slide := 0
	dp := make([][]int, m+1) // dp[i][j] represents the length of the square which lower right corner is located at (i, j).
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min2(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				if slide < dp[i][j] {
					slide = dp[i][j]
				}
			}
		}
	}

	return slide * slide
}

func min2(x, y, z int) int {
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
