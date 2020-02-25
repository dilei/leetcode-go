package dynamic_programming

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+sum {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
