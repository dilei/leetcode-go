package recursion

import (
	"fmt"
	"testing"
)

func Test_canPartitionKSubsets(t *testing.T) {
	res := canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4)
	fmt.Println(res)
}
