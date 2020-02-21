package backtracking

import (
	"strconv"
	"strings"
)

func getPermutation2(n int, k int) string {
	factorial := make(map[int]int, n+1)
	sum := 1
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		sum *= i
		factorial[i] = sum
	}

	numbers := []int{}
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}

	// base 0
	k--

	var strs []string
	for i := 1; i <= n; i++ {
		index := k / factorial[n-i]
		strs = append(strs, strconv.Itoa(numbers[index]))
		numbers = append(numbers[:index], numbers[index+1:]...)
		k -= index * factorial[n-i]
	}
	return strings.Join(strs, "")
}

// =====================================

func getPermutation(n int, k int) string {
	var res string
	var visited = make([]int, n)
	help60(n, factorial(n-1), k, &visited, &res)
	return res
}

func help60(n int, f int, k int, visited *[]int, res *string) {
	// 判断是第几组
	offset := k % f
	groupIndex := k / f
	if offset > 0 {
		groupIndex += 1
	}

	var i int
	for ; i < len(*visited) && groupIndex > 0; i++ {
		if (*visited)[i] == 0 {
			groupIndex--
		}
	}
	// 标记已使用
	(*visited)[i-1] = 1
	if n-1 > 0 {
		if offset == 0 {
			offset = f
		}
		*res += strconv.Itoa(i)
		help60(n-1, f/(n-1), offset, visited, res)
	} else {
		*res += strconv.Itoa(i)
	}
}

func factorial(n int) int {
	res := 1
	for i := n; i > 1; i-- {
		res *= i
	}
	return res
}
