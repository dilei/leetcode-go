package divide_conquer

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	global, local := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		local = max(nums[i], nums[i]+local)
		global = max(local, global)
	}

	return global
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
