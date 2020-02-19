package backtracking

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	help39(candidates, target, 0, &res, &cur)
	return res
}

func help39(data []int, target int, start int, res *[][]int, cur *[]int) {
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
		*cur = append(*cur, data[i])
		help39(data, target-data[i], i, res, cur)
		*cur = (*cur)[:len(*cur)-1] // 从下一个开始
	}
}
