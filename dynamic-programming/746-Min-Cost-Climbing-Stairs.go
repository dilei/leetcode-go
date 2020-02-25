package dynamic_programming

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	for i := 2; i < n; i++ {
		if cost[i-1] < cost[i-2] {
			cost[i] += cost[i-1]
		} else {
			cost[i] += cost[i-2]
		}
	}
	if cost[n-1] < cost[n-2] {
		return cost[n-1]
	}
	return cost[n-2]
}
