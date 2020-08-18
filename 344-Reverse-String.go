package leetcodego

import "fmt"

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func test() {
	s := []byte{'1', '2', '3'}
	fmt.Printf("%c\n", s)
	reverseString(s) // slice 传递底层的引用，会影响原来的s
	fmt.Printf("%c\n", s)
}
