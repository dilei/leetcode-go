package csnotes

import "fmt"

// 根据二叉树的前序遍历和中序遍历的结果，重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// preOrder: 3 9 20 15 7
// inOrder : 9 3 15 20 7

func rebuildBinaryTree(preOrder []int, inOrder []int) {
	root := new(BinaryTree)
	helpBuildTree(root, preOrder, inOrder, 0, 0, len(preOrder))
	fmt.Println("post order: ")
	postOrder(root)
	fmt.Println("pre order: ")
	preorder(root)
}

func preorder(root *BinaryTree) {
	if root != nil && root.Val > 0 {
		fmt.Print(root.Val)
		fmt.Println()
		preorder(root.Left)
		preorder(root.Right)
	}
}

func postOrder(root *BinaryTree) {
	if root != nil && root.Val > 0 {
		postOrder(root.Left)
		postOrder(root.Right)
		fmt.Print(root.Val)
		fmt.Println()
	}
}

func helpBuildTree(node *BinaryTree, preOrder []int, inOrder []int, preStart, inStart, length int) {
	var nodePos, leftLen, rightLen int

	if length <= 0 {
		node = nil
	} else {
		node.Val = preOrder[preStart]
		node.Left = &BinaryTree{}
		node.Right = &BinaryTree{}
		for pos, val := range inOrder {
			if val == node.Val {
				nodePos = pos
				break
			}
		}

		leftLen = nodePos - inStart
		rightLen = length - (leftLen + 1)
		helpBuildTree(node.Left, preOrder, inOrder, preStart+1, inStart, leftLen)
		helpBuildTree(node.Right, preOrder, inOrder, preStart+leftLen+1, nodePos+1, rightLen)
	}
}
