package divide_conquer

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	row := 0
	col := len(matrix[0]) - 1 // start from top right center
	for col >= 0 && row <= len(matrix[0])-1 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] { // target not in this row, because the line sorted asc from left to right
			row++
		} else if target < matrix[row][col] { // target not in this col, beacuse the col sorted asc from top to bottom
			col--
		}
	}
	return false
}
