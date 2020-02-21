package dynamic_programming

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

// 动态规划
func longestPalindrome2(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var res string
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			// j -i < 3 表示j和i之间之多有一个元素，也符合回文串
			// dp[i+1,j-1] ==true 表示i和j之间已经是回文串
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
