package leetcode_go

import "math"

func maxProduct(nums []int) int {
	r := float64(nums[0]) // 默认最大值
	for i, min, max := 1, r, r; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}

		max = math.Max(float64(nums[i]), float64(nums[i])*max)
		min = math.Min(float64(nums[i]), float64(nums[i])*min)

		r = math.Max(r, max)
	}
	return int(r)
}
