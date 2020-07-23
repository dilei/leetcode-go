package csnotes

import (
	"fmt"
	"testing"
)

func Test_entryNodeOfLoop(t *testing.T) {
	listNode := ListNode{1, nil}
	listNode2 := ListNode{2, nil}
	listNode3 := ListNode{3, nil}
	listNode4 := ListNode{4, nil}
	listNode5 := ListNode{5, nil}
	listNode6 := ListNode{6, nil}
	listNode7 := ListNode{7, nil}
	listNode8 := ListNode{8, nil}
	listNode.Next = &listNode2
	listNode2.Next = &listNode3
	listNode3.Next = &listNode4
	listNode4.Next = &listNode5
	listNode5.Next = &listNode6
	listNode6.Next = &listNode7
	listNode7.Next = &listNode8
	listNode8.Next = &listNode3
	res := entryNodeOfLoop(&listNode)
	if res == nil {
		fmt.Println("no loop")
	} else {
		fmt.Println(res.Val)
	}

}

func Test_deleteNode(t *testing.T) {
	listNode := ListNode{1, nil}
	listNode2 := ListNode{2, nil}
	listNode3 := ListNode{3, nil}
	listNode4 := ListNode{4, nil}
	listNode.Next = &listNode2
	listNode2.Next = &listNode3
	listNode3.Next = &listNode4

	listNode5 := ListNode{1, nil}
	// head
	deleteNode(&listNode, &listNode5)
	// end
	deleteNode(&listNode, &listNode4)
	// next
	deleteNode(&listNode, &listNode3)
}

func Test_print(t *testing.T) {
	print(2)
}

func Test_numOf1(t *testing.T) {
	cnt := numOf1(0x11)
	if cnt != 2 {
		t.Error(cnt)
	}
}

func Test_rectCover(t *testing.T) {
	res := rectCover(3)
	res2 := rectCover(5)
	fmt.Println(res, res2)
}

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

func Test_hasPath(t *testing.T) {
	chars := [][]byte{
		{'a', 'b', 't', 'g'},
		{'c', 'f', 'c', 's'},
		{'j', 'd', 'e', 'h'},
	}

	res := hasPath(chars, "bfce")
	res2 := hasPath(chars, "dddd")
	fmt.Println(res, res2)
}
