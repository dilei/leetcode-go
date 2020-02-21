package backtracking

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	res := solveNQueens(4)
	fmt.Println(res)
}
