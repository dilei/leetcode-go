package dynamic_programming

import "testing"

func Test_uniquePaths(t *testing.T) {
	res := uniquePaths(3, 2)
	if res != 3 {
		t.Error(res)
	}
}
