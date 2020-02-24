package dynamic_programming

func longestValidParentheses(s string) int {
	max := 0
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == ')' && len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			if len(stack) == 0 {
				if max < i {
					max = i
				}
			} else {
				if max < i-stack[len(stack)-1]-1 {
					max = i - stack[len(stack)-1] - 1
				}
			}
			stack = append(stack, i)
		}
	}

	if len(stack) == 0 {
		return len(s)
	} else if max < len(s)-stack[len(stack)-1]-1 {
		return len(s) - stack[len(stack)-1] - 1
	}
	return max
}

// 动态规划
func longestValidParentheses2(s string) int {
	var open, max int
	val := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		}
		if s[i] == ')' && open > 0 {
			val[i] = val[i-1] + 2
			// add prev exp：()(()) val[5] = val[4] + 2 + val[1]
			if i-val[i] > 0 {
				val[i] += val[i-val[i]]
			}
			open--
		}
		if max < val[i] {
			max = val[i]
		}
	}
	return max
}
