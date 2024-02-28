package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func toInts(l *ListNode) []int {
	var result []int
	for ; l != nil; l = l.Next {
		result = append(result, l.Val)
	}
	return result
}

func newList(x ...int) *ListNode {
	var l *ListNode
	for i := len(x) - 1; i >= 0; i-- {
		l = &ListNode{
			Val:  x[i],
			Next: l,
		}
	}
	return l
}

func swapPairs(head *ListNode) *ListNode {
	pfirst := &head
	for {
		first := *pfirst
		if first == nil {
			break
		}
		second := first.Next
		if second == nil {
			break
		}

		*pfirst = second

		first.Next = second.Next
		second.Next = first

		pfirst = &first.Next
	}

	return head
}
