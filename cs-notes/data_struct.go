package csnotes

// list node
type ListNode struct {
	Val  int
	Next *ListNode
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
