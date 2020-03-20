package dynamic_programming

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	// down to up
	high := len(triangle)
	minPath := triangle[high-1]
	for i := high - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ { // i is the len of this layer
			// the minPath of select j
			minPath[j] = min120(minPath[j], minPath[j+1]) + triangle[i][j]
		}
	}
	return minPath[0]
}

func min120(x, y int) int {
	if x < y {
		return x
	}
	return y
}
