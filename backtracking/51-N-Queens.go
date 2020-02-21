package backtracking

func solveNQueens(n int) [][]string {
	var res [][]string
	var cols = make([]int, n)     // 列是否占用
	var xhills = make([]int, 2*n) // 左上右下是否占用
	var yhills = make([]int, 2*n) // 左下右上是否占用
	var queens = make([]int, n)   // 存放n皇后的位置
	help51(0, n, &res, &cols, &xhills, &yhills, &queens)
	return res
}

func help51(row int, n int, res *[][]string, cols *[]int, xhills *[]int, yhills *[]int, queens *[]int) {
	if row >= n {
		return
	}
	for col := 0; col < n; col++ {
		if isOK(cols, xhills, yhills, n, row, col) {
			(*queens)[row] = col
			(*cols)[col] = 1
			(*xhills)[row-col+n-1] = 1
			(*yhills)[row+col] = 1
			if row == n-1 {
				addSolution(res, *queens, n)
			}
			help51(row+1, n, res, cols, xhills, yhills, queens)
			// 回溯,尝试下种可能
			(*queens)[row] = 0
			(*cols)[col] = 0
			(*xhills)[row-col+n-1] = 0
			(*yhills)[row+col] = 0
		}
	}
}

func addSolution(res *[][]string, queens []int, n int) {
	var tmp []string
	for i := 0; i < n; i++ {
		var str string
		col := queens[i]
		for j := 0; j < n; j++ {
			if j == col {
				str += "Q"
			} else {
				str += "."
			}
		}
		tmp = append(tmp, str)
	}
	*res = append(*res, tmp)
}

func isOK(cols *[]int, xhills *[]int, yhills *[]int, n int, row int, col int) bool {
	if (*cols)[col]+(*xhills)[row-col+n-1]+(*yhills)[row+col] > 0 {
		return false
	}
	return true
}
