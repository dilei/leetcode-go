package divide_conquer

import "strconv"

func addOperators(num string, target int) []string {
	var res []string
	if len(num) == 0 {
		return res
	}
	help282(&res, num, target, "", 0, 0, 0)
	return res
}

func help282(res *[]string, num string, target int, path string, pos int, evel int, mulit int) {
	if pos == len(num) {
		if target == evel {
			*res = append(*res, path)
		}
		return
	}

	for i := pos; i < len(num); i++ {
		// corner case: if current position is 0, we can only use it as a single digit number, should be 0
		// if it is not a single digit number with leading 0, it should be considered as an invalid number
		if pos != i && num[pos] == '0' { // num start with  0 is error num
			break
		}

		cur, _ := strconv.Atoi(num[pos : i+1])
		if pos == 0 {
			help282(res, num, target, path+strconv.Itoa(cur), i+1, cur, cur)
		} else {
			help282(res, num, target, path+"+"+strconv.Itoa(cur), i+1, evel+cur, cur)
			help282(res, num, target, path+"-"+strconv.Itoa(cur), i+1, evel-cur, -cur)
			help282(res, num, target, path+"*"+strconv.Itoa(cur), i+1, evel-mulit+mulit*cur, mulit*cur) // 乘法需要重新结合
		}
	}
}
