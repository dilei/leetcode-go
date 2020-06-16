package csnotes

import (
	"fmt"
	"testing"
)

func Test_duplicate(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5}
	res, ok := duplicate(nums)
	fmt.Println(res, ok)
}

func Test_find(t *testing.T) {
	nums := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	res := find(18, nums)
	res1 := find(17, nums)
	fmt.Println(res, res1)
}

func Test_printList(t *testing.T) {
	listNode := ListNode{1, nil}
	listNode2 := ListNode{2, nil}
	listNode3 := ListNode{3, nil}
	listNode.next = &listNode2
	listNode2.next = &listNode3
	node := &listNode
	for node != nil {
		fmt.Print(node.val)
		fmt.Print("\t")
		node = node.next
	}
	fmt.Println()

	printList(&listNode)
	fmt.Println()
}

func Test_printList2(t *testing.T) {
	listNode := ListNode{1, nil}
	listNode2 := ListNode{2, nil}
	listNode3 := ListNode{3, nil}
	listNode.next = &listNode2
	listNode2.next = &listNode3

	fmt.Println("before: ")
	node := &listNode
	for node != nil {
		fmt.Print(node.val)
		fmt.Print("\t")
		node = node.next
	}
	fmt.Println()

	res := printList2(&listNode)

	fmt.Println("after: ")
	node = res
	for node != nil {
		fmt.Print(node.val)
		fmt.Print("\t")
		node = node.next
	}
	fmt.Println()
}
