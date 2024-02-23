package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(nums ...int) *ListNode {
	var h *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		h = &ListNode{
			Next: h,
			Val:  nums[i],
		}
	}
	return h
}

func toInts(h *ListNode) []int {
	var result []int
	for ; h != nil; h = h.Next {
		result = append(result, h.Val)
	}
	return result
}

func rotateRight(head *ListNode, k int) *ListNode {
	count := 0
	var last *ListNode
	for h := head; h != nil; h = h.Next {
		count++
		last = h
	}
	if count == 0 {
		return head
	}

	shift := k % count
	if shift == 0 {
		return head
	}

	result := head
	ptr := &result
	for i := 0; i < count-shift; i++ {
		ptr = &(*ptr).Next
	}

	last.Next = head
	result = *ptr
	*ptr = nil

	return result
}
