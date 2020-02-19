package backtracking

import (
	"sort"
)

func combinationSumII(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var cur []int
	help40(candidates, target, 0, &res, &cur)
	return res
}

func help40(data []int, target int, start int, res *[][]int, cur *[]int) {
	if target == 0 {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	if target < 0 {
		return
	}

	for i := start; i < len(data); i++ {
		if i > start && data[i] == data[i-1] { // 因为已经排序了
			continue
		}
		*cur = append(*cur, data[i])
		help40(data, target-data[i], i+1, res, cur) // 不能重复使用，从i+1尝试
		*cur = (*cur)[:len(*cur)-1]                 // 从下一个开始
	}
}
