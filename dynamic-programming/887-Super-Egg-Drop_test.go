package dynamic_programming

import (
	"fmt"
	"testing"
)

func Test_superEggDrop(t *testing.T) {
	res := superEggDrop(1, 2)
	res2 := superEggDrop(2, 6)
	res3 := superEggDrop(3, 14)
	fmt.Println(res, res2, res3)
}
