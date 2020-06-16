package csnotes

// 给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。
//
// Consider the following matrix:
// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
//
// Given target = 5, return true.
// Given target = 20, return false.

func find(target int, nums [][]int) bool {
	if len(nums) == 0 || len(nums[0]) == 0 {
		return false
	}

	// nums[row][col] 从右上角开始找
	row := 0
	col := len(nums[0]) - 1

	for col > 0 && row < len(nums) {
		if nums[row][col] == target {
			return true
		} else if nums[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}