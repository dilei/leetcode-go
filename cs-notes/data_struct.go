package csnotes

import "fmt"

// list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// 含父节点
type BinaryTree2 struct {
	Val   int
	Left  *BinaryTree2
	Right *BinaryTree2
	Next  *BinaryTree2 // 父节点
}
