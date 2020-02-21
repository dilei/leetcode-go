package backtracking

func totalNQueens(n int) int {
	var res int
	var cols = make([]int, n)     // 列是否占用
	var xhills = make([]int, 2*n) // 左上右下是否占用
	var yhills = make([]int, 2*n) // 左下右上是否占用
	var queens = make([]int, n)   // 存放n皇后的位置
	help52(0, n, &res, &cols, &xhills, &yhills, &queens)
	return res
}

func help52(row int, n int, res *int, cols *[]int, xhills *[]int, yhills *[]int, queens *[]int) {
	if row >= n {
		return
	}
	for col := 0; col < n; col++ {
		if isOK2(cols, xhills, yhills, n, row, col) {
			// 主对角线上: 行 - 列 = 常数  "\"
			// 副对角线上：行 + 列 = 常数  "/"
			(*queens)[row] = col
			(*cols)[col] = 1
			(*xhills)[row-col] = 1
			(*yhills)[row+col] = 1
			if row == n-1 {
				*res++
			}
			help52(row+1, n, res, cols, xhills, yhills, queens)
			// 回溯,尝试下种可能
			(*queens)[row] = 0
			(*cols)[col] = 0
			(*xhills)[row-col] = 0
			(*yhills)[row+col] = 0
		}
	}
}

func isOK2(cols *[]int, xhills *[]int, yhills *[]int, n int, row int, col int) bool {
	if (*cols)[col]+(*xhills)[row-col]+(*yhills)[row+col] > 0 {
		return false
	}
	return true
}
