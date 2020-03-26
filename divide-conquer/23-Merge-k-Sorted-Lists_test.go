package divide_conquer

import (
	"fmt"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l6 := &ListNode{Val: 6}
	l7 := &ListNode{Val: 1}
	l8 := &ListNode{Val: 4}
	l1.Next = l4
	l4.Next = l5

	l7.Next = l3
	l3.Next = l8

	l2.Next = l6

	l := []*ListNode{
		l1,
		l7,
		l2,
	}
	res := mergeKLists(l)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
	fmt.Println()
}
