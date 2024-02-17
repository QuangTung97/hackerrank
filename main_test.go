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
	v := reverseLinkedList(toList(1, 2, 3), 4)
	assert.Equal(t, []int{3, 2, 1}, fromList(v))

	v = reverseLinkedList(toList(1, 2, 3, 4, 5), 3)
	assert.Equal(t, []int{3, 2, 1, 4, 5}, fromList(v))
}

func TestFindK(t *testing.T) {
	l := toList(1, 2, 3, 4, 5)
	ok := findNextK(&l, 2)
	assert.Equal(t, true, ok)

	l = toList(1, 2, 3, 4, 5)
	ok = findNextK(&l, 5)
	assert.Equal(t, true, ok)

	l = toList(1, 2, 3, 4, 5)
	ok = findNextK(&l, 6)
	assert.Equal(t, false, ok)
}

func TestExample_KGroup(t *testing.T) {
	r := reverseKGroup(toList(1, 2, 3, 4, 5), 2)
	assert.Equal(t, []int{2, 1, 4, 3, 5}, fromList(r))

	r = reverseKGroup(toList(1, 2, 3, 4, 5), 3)
	assert.Equal(t, []int{3, 2, 1, 4, 5}, fromList(r))
}
