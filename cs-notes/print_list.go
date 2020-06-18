package csnotes

import "fmt"

// 从尾到头反过来打印出每个结点的值。

// 递归
func printList(listNode *ListNode) {
	if listNode != nil && listNode.Next != nil {
		printList(listNode.Next)
	}
	fmt.Print(listNode.Val)
	fmt.Print("\t")
}

// 指针反转
func printList2(listNode *ListNode) *ListNode {
	head := new(ListNode)
	node := listNode
	for node != nil {
		tmp := node.Next
		node.Next = head.Next
		head.Next = node
		node = tmp
	}

	return head.Next
}
