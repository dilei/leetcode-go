package leetcodego

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var stack []byte
	m := map[byte]byte{
		91:  93,
		123: 125,
		40:  41,
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 {
			top := stack[len(stack)-1]
			if m[top] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else { // "]"
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
