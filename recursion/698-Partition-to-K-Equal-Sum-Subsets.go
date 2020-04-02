package recursion

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if k <= 0 || sum%k != 0 {
		return false
	}
	visited := make([]int, len(nums))
	return help698(nums, visited, 0, k, 0, 0, sum/k)
}

func help698(nums []int, visited []int, start int, k int, cur_sum int, cur_num int, target int) bool {
	if k == 1 {
		return true
	}
	if cur_sum == target && cur_num > 0 {
		return help698(nums, visited, 0, k-1, 0, 0, target)
	}
	for i := start; i < len(nums); i++ {
		if visited[i] == 0 {
			visited[i] = 1
			if help698(nums, visited, i+1, k, cur_sum+nums[i], cur_num+1, target) {
				return true
			}
			visited[i] = 0
		}
	}
	return false
}
