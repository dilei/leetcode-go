package csnotes

func entryNodeOfLoop(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast := head
	slow := head
	for {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		if fast == slow {
			break
		}
	}
	fast = head
	for {
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	return slow
}
