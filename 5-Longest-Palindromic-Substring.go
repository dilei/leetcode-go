package leetcode_go

// 全局变量可能会有并发问题
// var low, maxlen int

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var low, maxlen int
	for i := 0; i < len(s)-1; i++ {
		help5(s, i, i+1, &low, &maxlen) // 偶数位
		help5(s, i, i, &low, &maxlen)   // 奇数位
	}
	return s[low : low+maxlen]
}

func help5(s string, j, k int, low, maxlen *int) {
	for j >= 0 && k < len(s) && s[j] == s[k] {
		j--
		k++
	}
	if *maxlen < k-j-1 {
		*maxlen = k - j - 1
		*low = j + 1 // 加一是因为上一步向前移了一位
	}
}
