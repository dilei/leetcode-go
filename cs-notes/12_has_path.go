package csnotes

import (
	"fmt"
)

func hasPath(chars [][]byte, str string) bool {
	if len(chars) == 0 || len(chars[0]) == 0 {
		return false
	}

	rows := len(chars)
	cols := len(chars[0])
	var used [][]bool
	for range chars {
		used = append(used, make([]bool, cols))
	}
	fmt.Println(used)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if help12(chars, used, rows, cols, i, j, 0, str) {
				return true
			}
		}
	}
	return false
}

func help12(chars [][]byte, used [][]bool, rows, cols, row, col, pathLen int, str string) bool {
	if pathLen == len(str) {
		return true
	}

	if row < 0 || row >= rows || col < 0 || col >= cols || used[row][col] || chars[row][col] != str[pathLen] {
		return false
	}

	path := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	used[row][col] = true
	for _, v := range path {
		if help12(chars, used, rows, cols, row+v[0], col+v[1], pathLen+1, str) {
			return true
		}
	}
	used[row][col] = false
	return false
}
