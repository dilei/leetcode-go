package leetcode_go

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}

	if p[1] != '*' {
		if len(s) == 0 {
			return false
		}
		return (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	// 处理相同前缀且含有*的情况 exp: s: aa p: a*
	for len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
		if isMatch(s, p[2:]) {
			return true
		}
		s = s[1:]
	}

	return isMatch(s, p[2:])
}
