package dynamic_programming

import "math"

func maxProfit(prices []int) int {
	maxCur, maxAll := 0, 0
	for i := 1; i < len(prices); i++ {
		maxCur += prices[i] - prices[i-1]
		maxCur = max(0, maxCur)
		maxAll = max(maxCur, maxAll)
	}
	return maxAll
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit2(prices []int) int {
	min := math.MaxInt32
	maxPrice := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > maxPrice {
				maxPrice = prices[i] - min
			}
		}
	}
	return maxPrice
}
