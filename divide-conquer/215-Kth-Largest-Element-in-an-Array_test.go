package divide_conquer

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	res := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	fmt.Println(res)
}
