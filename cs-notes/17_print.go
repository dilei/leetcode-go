package csnotes

import (
	"fmt"
)

// 打印从 1 到最大的 n 位数

func print(n int) {
	if n <= 0 {
		return
	}
	// fmt.Println(byte(100))
	// fmt.Println(string(100))
	number := make([]int, n)
	print1ToMaxOfNDigits(number, 0)
}

func print1ToMaxOfNDigits(number []int, digit int) {
	if digit == len(number) {
		printNumber(number)
		return
	}

	for i := 0; i < 10; i++ {
		number[digit] = i
		print1ToMaxOfNDigits(number, digit+1)
	}
}

func printNumber(number []int) {
	index := 0
	for index < len(number) && number[index] == 0 {
		index++
	}
	for index < len(number) {
		fmt.Print(number[index])
		index++
	}
	fmt.Println()
}
