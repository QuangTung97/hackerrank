package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(num ...int) *ListNode {
	var h *ListNode
	for i := len(num) - 1; i >= 0; i-- {
		h = &ListNode{
			Next: h,
			Val:  num[i],
		}
	}
	return h
}

func toInts(h *ListNode) []int {
	var res []int
	for e := h; e != nil; e = e.Next {
		res = append(res, e.Val)
	}
	return res
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := head
	p := &result
	right := head
	space := 0

	for ; right != nil; right = right.Next {
		space++
		if space > n {
			p = &(*p).Next
			space--
		}
	}

	if space == n {
		*p = (*p).Next
	}

	return result
}
