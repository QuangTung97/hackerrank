package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func toList(arr ...int) *ListNode {
	var result *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		result = &ListNode{
			Val:  arr[i],
			Next: result,
		}
	}
	return result
}

func fromList(a *ListNode) []int {
	var result []int
	for a != nil {
		result = append(result, a.Val)
		a = a.Next
	}
	return result
}

func TestExample(t *testing.T) {
	x := toList(1, 2, 3)
	assert.Equal(t, 1, x.Val)
	assert.Equal(t, 2, x.Next.Val)
	assert.Equal(t, 3, x.Next.Next.Val)
	assert.Equal(t, []int{1, 2, 3}, fromList(x))
}

func TestAlgo(t *testing.T) {
	r := mergeTwoLists(
		toList(1, 2, 4),
		toList(1, 3, 4),
	)
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, fromList(r))
}

func TestAlgo_Empty(t *testing.T) {
	r := mergeTwoLists(
		toList(),
		toList(),
	)
	assert.Equal(t, []int(nil), fromList(r))
}

func TestAlgo_Single(t *testing.T) {
	r := mergeTwoLists(
		toList(),
		toList(2),
	)
	assert.Equal(t, []int{2}, fromList(r))
}
