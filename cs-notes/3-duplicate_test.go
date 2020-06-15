package cs-notes

import "testing"

func Test_duplicate(t *testing.T)  {
	nums := []int{2, 3, 1, 0, 2, 5}
	res, ok := duplicate(nums)
	fmt.Println(res, ok)
}