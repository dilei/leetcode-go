package backtracking

func solveSudoku(board [][]byte) {
	help37(board)
}

var chars = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

/*
func help37(board *[][]byte) bool {
	for i := 0; i < len(*board); i++ { // 列
		for j := 0; j < len((*board)[0]); j++ { // 行
			if (*board)[i][j] == '.' {
				for _, c := range chars {
					if isValid(*board, i, j, c) {
						(*board)[i][j] = c
						if help37(board) {
							return true
						} else {
							(*board)[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}
*/

func help37(board [][]byte) bool {
	for i := 0; i < len(board); i++ { // 列
		for j := 0; j < len(board[0]); j++ { // 行
			if board[i][j] == '.' {
				for _, c := range chars {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if help37(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	regionStartRow := 3 * (row / 3)
	regionStartCol := 3 * (col / 3)
	for i := 0; i < 9; i++ {
		if board[i][col] == c { // 检查列
			return false
		}
		if board[row][i] == c { // 检查行
			return false
		}
		if board[regionStartRow+i/3][regionStartCol+i%3] == c { // 检查3x3 控制走过9个小格子
			return false
		}
	}
	return true
}
