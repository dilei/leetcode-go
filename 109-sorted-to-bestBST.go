package leetcodego

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil

	treeNode := &TreeNode{Val: slow.Val}
	treeNode.Left = sortedListToBST(head)
	treeNode.Right = sortedListToBST(slow.Next)
	return treeNode
}
