package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	a := list1
	b := list2

	if a == nil && b == nil {
		return nil
	}

	if a == nil {
		return &ListNode{
			Val:  b.Val,
			Next: mergeTwoLists(a, b.Next),
		}
	}

	if b == nil || a.Val < b.Val {
		return &ListNode{
			Val:  a.Val,
			Next: mergeTwoLists(a.Next, b),
		}
	}

	return &ListNode{
		Val:  b.Val,
		Next: mergeTwoLists(a, b.Next),
	}
}
