package dynamic_programming

import (
	"math"
)

func maxProfit123(prices []int) int {
	oneBuyOneSell := 0
	twoBuyTwoSell := 0
	oneBuy := math.MaxInt32
	twoBuy := math.MaxInt32

	for _, p := range prices {
		oneBuy = min123(oneBuy, -p) // 当天价格为负，表示买入
		oneBuyOneSell = max(oneBuyOneSell, p+oneBuy)
		twoBuy = min123(twoBuy, oneBuyOneSell-p)
		twoBuyTwoSell = max(twoBuyTwoSell, p+twoBuy)
	}
	return twoBuyTwoSell
}

func min123(x, y int) int {
	if x < y {
		return x
	}
	return y
}
