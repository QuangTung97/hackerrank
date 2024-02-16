package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := addTwoNumbersRemainder(l1, l2, 0)
	if result == nil {
		return &ListNode{
			Val: 0,
		}
	}
	return result
}

func addTwoNumbersRemainder(l1 *ListNode, l2 *ListNode, remainder int) *ListNode {
	aVal := 0
	var nextL1 *ListNode
	if l1 != nil {
		aVal = l1.Val
		nextL1 = l1.Next
	}

	bVal := 0
	var nextL2 *ListNode
	if l2 != nil {
		bVal = l2.Val
		nextL2 = l2.Next
	}

	sum := aVal + bVal + remainder
	if sum == 0 && l1 == nil && l2 == nil {
		return nil
	}

	newRemainder := 0
	if sum >= 10 {
		newRemainder = 1
		sum -= 10
	}

	return &ListNode{
		Val:  sum,
		Next: addTwoNumbersRemainder(nextL1, nextL2, newRemainder),
	}
}
