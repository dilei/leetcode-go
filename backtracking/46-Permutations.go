package backtracking

func permute(nums []int) [][]int {
	var res [][]int
	var cur []int
	help46(nums, &res, &cur)
	return res
}
func isUse(val int, cur []int) bool {
	for _, j := range cur {
		if j == val {
			return true
		}
	}
	return false
}

func help46(nums []int, res *[][]int, cur *[]int) {
	if len(*cur) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if isUse(nums[i], *cur) {
			continue
		}

		*cur = append(*cur, nums[i])
		help46(nums, res, cur)
		*cur = (*cur)[:len(*cur)-1]
	}
}
