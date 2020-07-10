package csnotes

// 二进制中 1 的个数
// n & (n-1)
// java: Integer.bitCount(n)

func numOf1(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		n &= n - 1
	}
	return cnt
}
