package csnotes

// 矩形覆盖
// c
// 我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形，总共有多少种方法？

func rectCover(n int) int {
	if n <= 2 {
		return n
	}

	pre1 := 1
	pre2 := 2
	var result int
	for i:=3;  i<=n; i++ {
		result = pre1 + pre2
		pre1 = pre2
		pre2 = result
	}

	return result
}