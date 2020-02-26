package dynamic_programming

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	maxArea := 0
	hight := make([]int, n) // hight[j]表示到j的高度
	left := make([]int, n)  // left[j] 表示到j的左端开始节点
	right := make([]int, n) // rifht[j] 表示到j的右端开始节点
	// 所以j点围成的面积就是到j的（right - left） * heght

	for i := range right {
		right[i] = n // 填充默认值
	}
	for i := 0; i < m; i++ {
		cur_left := 0
		cur_right := n
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				// left
				if left[j] < cur_left {
					left[j] = cur_left
				}

				// hight
				hight[j] += 1

			} else {
				left[j] = 0
				cur_left = j + 1

				hight[j] = 0
			}
		}
		for j := n - 1; j >= 0; j-- {
			// right
			if matrix[i][j] == '1' {
				if right[j] > cur_right {
					right[j] = cur_right
				}
			} else {
				right[j] = n
				cur_right = j
			}

		}
		for j := 0; j < n; j++ {
			if maxArea < (right[j]-left[j])*hight[j] {
				maxArea = (right[j] - left[j]) * hight[j]
			}
		}
	}

	return maxArea
}
