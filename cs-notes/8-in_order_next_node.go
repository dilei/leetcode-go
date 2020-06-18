package csnotes

// 给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回 。
// 注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。

func inOrderNextNode(tree *BinaryTree2) int {
	if tree.Right != nil { // 存在右节点，那么下一个节点就是右节点的最左节点
		node := tree.Right
		for node.Left != nil {
			node = node.Left
		}
		return node.Val
	}
	// 不存在右节点，那么下一个节点就是向上找左节点是此节点祖先节点的那个节点
	for tree.Next != nil {
		parent := tree.Next
		if parent.Left == tree {
			return parent.Val
		}
		tree = tree.Next
	}
	return 0
}