package dynamic_programming

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	letters := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		letters[s1[i]-'a']++
		letters[s2[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if letters[i] > 0 { // 含有不同字符
			return false
		}
	}

	// 符合题意的字符串，中间会存在一个分割点，使得前后俩部分和原始字符串前后两部分相同，或者反转后相同
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}
