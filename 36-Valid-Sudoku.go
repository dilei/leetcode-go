package leetcode_go

import "strconv"

func isValidSudoku(board [][]byte) bool {
	row := make(map[string]bool, len(board))
	col := make(map[string]bool, len(board))
	block := make(map[string]bool, len(board))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				num := board[i][j]
				rowKey := strconv.Itoa(int(num-'0')) + " row " + strconv.Itoa(i)
				colKey := strconv.Itoa(int(num-'0')) + " col " + strconv.Itoa(j)
				blockKey := strconv.Itoa(int(num-'0')) + " block " + strconv.Itoa(i/3) + strconv.Itoa(j/3)
				if row[rowKey] || col[colKey] || block[blockKey] {
					return false
				}
				row[rowKey] = true
				col[colKey] = true
				block[blockKey] = true
			}
		}
	}

	return true
}

func isValidSudoku2(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' && isOK(i, j, board) == false {
				return false
			}
		}
	}

	return true
}

func isOK(i, j int, board [][]byte) bool {
	row := 3 * (i / 3)
	col := 3 * (j / 3)
	c := board[i][j]
	board[i][j] = '.'
	defer func() {
		board[i][j] = c
	}()
	for k := 0; k < 9; k++ {
		if board[k][j] == c {
			return false
		}
		if board[i][k] == c {
			return false
		}
		if board[row+k/3][col+k%3] == c {
			return false
		}
	}
	return true
}
