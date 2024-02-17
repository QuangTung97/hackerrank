package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func toList(a ...int) *ListNode {
	var result *ListNode
	for i := len(a) - 1; i >= 0; i-- {
		result = &ListNode{
			Val:  a[i],
			Next: result,
		}
	}
	return result
}

func fromList(l *ListNode) []int {
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}

func TestToList(t *testing.T) {
	l := toList(1, 2, 3)
	assert.Equal(t, []int{1, 2, 3}, fromList(l))
}

func TestExample(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		r := mergeKLists([]*ListNode{
			toList(2, 6),
			toList(1, 4, 5),
			toList(1, 3, 4),
		})
		assert.Equal(t, []int{1, 1, 2, 3, 4, 4, 5, 6}, fromList(r))
	})
}
