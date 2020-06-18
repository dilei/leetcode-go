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
	listNode.Next = &listNode2
	listNode2.Next = &listNode3
	node := &listNode
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print("\t")
		node = node.Next
	}
	fmt.Println()

	printList(&listNode)
	fmt.Println()
}

func Test_printList2(t *testing.T) {
	listNode := ListNode{1, nil}
	listNode2 := ListNode{2, nil}
	listNode3 := ListNode{3, nil}
	listNode.Next = &listNode2
	listNode2.Next = &listNode3

	fmt.Println("before: ")
	node := &listNode
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print("\t")
		node = node.Next
	}
	fmt.Println()

	res := printList2(&listNode)

	fmt.Println("after: ")
	node = res
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print("\t")
		node = node.Next
	}
	fmt.Println()
}

func Test_rebuildBinaryTree(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	rebuildBinaryTree(preOrder, inOrder)
}

func Test_inOrderNextNode(t *testing.T) {
	//               6
	//      2                 7
	// 1         4
	//      3         5
	root := &BinaryTree2{Val: 6}
	node1 := &BinaryTree2{Val: 2}
	node2 := &BinaryTree2{Val: 7}
	node3 := &BinaryTree2{Val: 1}
	node4 := &BinaryTree2{Val: 4}
	node5 := &BinaryTree2{Val: 3}
	node6 := &BinaryTree2{Val: 5}
	root.Left = node1
	root.Right = node2
	node2.Next = root
	node1.Next = root

	node1.Left = node3
	node1.Right = node4
	node3.Next = node1
	node4.Next = node1

	node4.Left = node5
	node4.Right = node6
	node5.Next = node4
	node6.Next = node4

	res := inOrderNextNode(node1)
	res2 := inOrderNextNode(node6)
	fmt.Println(res, res2)
}
