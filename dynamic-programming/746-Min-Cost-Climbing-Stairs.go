package dynamic_programming

// you can either start from the step with index 0, or the step with index 1
// ==> cost[0] = cost[0]
// ==> cost[1] = cost[1]
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	for i := 2; i < n; i++ {
		if cost[i-1] < cost[i-2] {
			cost[i] += cost[i-1]
		} else {
			cost[i] += cost[i-2]
		}
	}
	// 走到n-1或n-2都可以直达top
	if cost[n-1] < cost[n-2] {
		return cost[n-1]
	}
	return cost[n-2]
}
