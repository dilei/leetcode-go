package csnotes

import "fmt"

// 从尾到头反过来打印出每个结点的值。

// 递归
func printList(listNode *ListNode) {
	if listNode != nil && listNode.next != nil {
		printList(listNode.next)
	}
	fmt.Print(listNode.val)
	fmt.Print("\t")
}

// 指针反转
func printList2(listNode *ListNode) *ListNode {
	head := new(ListNode)
	node := listNode
	for node != nil {
		tmp := node.next
		node.next = head.next
		head.next = node
		node = tmp
	}

	return head.next
}
