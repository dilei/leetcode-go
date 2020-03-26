package divide_conquer

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return help23(lists, 0, len(lists)-1)
}

func help23(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if start > end {
		return nil
	}

	// mid := (end + start) / 2
	mid := (end-start)/2 + start
	l1 := help23(lists, start, mid)
	l2 := help23(lists, mid+1, end)

	h := &ListNode{}
	cur := h
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return h.Next
}
