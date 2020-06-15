package cs-notes

// 在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
// 
// Input:
// {2, 3, 1, 0, 2, 5}
// 
// Output:
// 2

func duplicate(nums []int) (int, bool)  {
	l := len(nums)
	if l == 0 {
		return false, 0
	}

	for i:=0; i<l; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] { // find
				return nums[i], true
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return 0, false
}