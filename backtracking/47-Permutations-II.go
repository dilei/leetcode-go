package backtracking

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var cur []int
	help47(nums, map[int]bool{}, &res, &cur)
	return res
}

func help47(nums []int, used map[int]bool, res *[][]int, cur *[]int) {
	if len(*cur) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && used[i-1] {
			continue
		}
		*cur = append(*cur, nums[i])
		used[i] = true
		help47(nums, used, res, cur)
		used[i] = false
		*cur = (*cur)[:len(*cur)-1]
	}
}
