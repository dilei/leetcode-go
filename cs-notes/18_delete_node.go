package csnotes

// O(1)时间 删除list节点

func deleteNode(head, deleteNode *ListNode) {
	if head == nil || deleteNode == nil {
		return
	}

	if deleteNode.Next != nil {
		tmpNode := deleteNode.Next
		deleteNode.Val = tmpNode.Val
		deleteNode.Next = tmpNode.Next
	} else { // 尾节点
		if head == deleteNode { // 说明只有以这个节点
			head = nil
		} else {
			cur := head
			for cur != nil && cur.Next != deleteNode {
				cur = cur.Next
			}
			if cur != nil {
				cur.Next = nil
			}
		}
	}

	PrintListNode(head)
}
