package divide_conquer

import "strconv"

func diffWaysToCompute(input string) []int {
	var res []int
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			part1 := input[:i]
			part2 := input[i+1:]
			res1 := diffWaysToCompute(part1)
			res2 := diffWaysToCompute(part2)
			for _, r := range res1 {
				for _, r2 := range res2 {
					c := 0
					switch input[i] {
					case '+':
						c = r + r2
					case '-':
						c = r - r2
					case '*':
						c = r * r2
					}
					res = append(res, c)
				}
			}
		}
	}
	if len(res) == 0 {
		c, _ := strconv.Atoi(input)
		res = append(res, c)
	}
	return res
}
