package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func drain(it *BSTIterator) []int {
	var res []int
	for {
		res = append(res, it.Next())
		if !it.HasNext() {
			return res
		}
	}
}

func TestCount(t *testing.T) {
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 7}
	n3 := &TreeNode{Val: 9}
	n4 := &TreeNode{Val: 15}
	n5 := &TreeNode{Val: 20}

	n2.Left = n1
	n2.Right = n4
	n4.Left = n3
	n4.Right = n5

	p := getAllLeftNodes(n2)
	assert.Equal(t, []int{7, 3}, pathValues(p))
	assert.Equal(t, []int{15, 9}, pathValues(getAllLeftNodes(n2.Right)))

	// check next path
	p = getAllLeftNodes(n2)

	p1 := computeNextPath(p)
	assert.Equal(t, []int{7}, pathValues(p1))

	p2 := computeNextPath(p1)
	assert.Equal(t, []int{7, 15, 9}, pathValues(p2))

	p3 := computeNextPath(p2)
	assert.Equal(t, []int{7, 15}, pathValues(p3))

	p4 := computeNextPath(p3)
	assert.Equal(t, []int{7, 15, 20}, pathValues(p4))

	p5 := computeNextPath(p4)
	assert.Equal(t, []int(nil), pathValues(p5))

	// check drain it
	it := Constructor(n2)
	nodes := drain(&it)
	assert.Equal(t, []int{3, 7, 9, 15, 20}, nodes)
}
