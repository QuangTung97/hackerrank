package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type priorityQueue struct {
	data []*ListNode
}

var _ heap.Interface = &priorityQueue{}

func (q *priorityQueue) Len() int {
	return len(q.data)
}

func (q *priorityQueue) Less(i, j int) bool {
	return q.data[i].Val < q.data[j].Val
}

func (q *priorityQueue) Swap(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *priorityQueue) Push(x any) {
	q.data = append(q.data, x.(*ListNode))
}

func (q *priorityQueue) Pop() any {
	n := len(q.data)
	r := q.data[n-1]
	q.data = q.data[:n-1]
	return r
}

func mergeKLists(input []*ListNode) *ListNode {
	lists := make([]*ListNode, 0, len(input))
	for _, e := range input {
		if e == nil {
			continue
		}
		lists = append(lists, e)
	}

	q := &priorityQueue{
		data: lists,
	}
	heap.Init(q)

	var head *ListNode
	next := &head

	for len(q.data) > 0 {
		x := heap.Pop(q).(*ListNode)

		newElem := x.Next
		if newElem != nil {
			heap.Push(q, newElem)
		}
		x.Next = nil

		*next = x
		next = &x.Next
	}

	return head
}
