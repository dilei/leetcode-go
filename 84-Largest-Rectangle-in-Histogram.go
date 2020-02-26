package leetcode_go

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	h := make([]int, n+1)
	copy(h, heights)
	max, i := 0, 0
	var stack []int
	for i < len(h) {
		if len(stack) == 0 || h[stack[len(stack)-1]] < h[i] {
			stack = append(stack, i)
			i++
		} else {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := -1
			if len(stack) == 0 {
				tmp = h[t] * i
			} else {
				tmp = h[t] * (i - stack[len(stack)-1] - 1)
			}
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
