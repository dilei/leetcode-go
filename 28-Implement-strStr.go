package leetcodego

func strStr(haystack string, needle string) int {
	// return strings.Index(haystack, needle)

	// 2
	for i := 0; ; i++ { // len of haystack
		for j := 0; ; j++ { // lend of needle
			if j == len(needle) { // find the str of needle
				return i
			}
			if i+j == len(haystack) { // needle has len of left, but len of haystack is done
				return -1
			}
			if haystack[i+j] != needle[j] { // not match
				break
			}
		}
	}
}
