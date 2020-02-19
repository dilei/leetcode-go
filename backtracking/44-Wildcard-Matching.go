package backtracking

func isMatch2(s string, p string) bool {
	var si, pi, match int
	var startIndex = -1
	for si < len(s) {
		if pi < len(p) && (p[pi] == s[si] || p[pi] == '?') {
			si++
			pi++
		} else if pi < len(p) && p[pi] == '*' {
			startIndex = pi
			match = si
			pi++
		} else if startIndex != -1 {
			pi = startIndex + 1
			match++
			si = match
		} else {
			return false
		}
	}

	for pi < len(p) && p[pi] == '*' {
		pi++
	}
	return pi == len(p)
}
