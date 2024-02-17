package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(l *ListNode, k int) *ListNode {
	reverseLinkedListHead(&l, k)
	return l
}

func reverseLinkedListHead(head **ListNode, k int) {
	node := (*head).Next
	prev := *head

	num := k - 1
	for i := 0; i < num && node != nil; i++ {
		next := node.Next
		node.Next = *head
		prev.Next = next

		*head = node
		node = next
	}
}

func findNextK(head **ListNode, k int) bool {
	for i := 0; i < k; i++ {
		if *head == nil {
			return false
		}
		head = &(*head).Next
	}
	return true
}

func reverseKGroup(inputHead *ListNode, k int) *ListNode {
	currentHead := &inputHead
	for {
		ok := findNextK(currentHead, k)
		if !ok {
			return inputHead
		}
		nextHead := &(*currentHead).Next
		reverseLinkedListHead(currentHead, k)
		currentHead = nextHead
	}
}
