package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func toList(elems ...int) *ListNode {
	var result *ListNode
	next := &result
	for _, x := range elems {
		*next = &ListNode{
			Val: x,
		}
		next = &(*next).Next
	}
	return result
}

func TestTwoNumbers_ToList(t *testing.T) {
	l := toList(3, 2, 1)
	assert.Equal(t, 3, l.Val)
	assert.Equal(t, 2, l.Next.Val)
	assert.Equal(t, 1, l.Next.Next.Val)
	assert.Nil(t, l.Next.Next.Next)
}

func assertEqual(t *testing.T, expect []int, l *ListNode) {
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	assert.Equal(t, expect, result)
}

func TestTwoNumbers(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		result := addTwoNumbers(
			toList(2),
			toList(3),
		)
		assertEqual(t, []int{5}, result)
	})

	t.Run("normal", func(t *testing.T) {
		result := addTwoNumbers(
			toList(6),
			toList(7),
		)
		assertEqual(t, []int{3, 1}, result)
	})

	t.Run("add two zeros", func(t *testing.T) {
		result := addTwoNumbers(
			toList(0),
			toList(0),
		)
		assertEqual(t, []int{0}, result)
	})

	t.Run("one longer than other", func(t *testing.T) {
		result := addTwoNumbers(
			toList(1, 2, 3),
			toList(7),
		)
		assertEqual(t, []int{8, 2, 3}, result)
	})

	t.Run("one longer than other, with remainders", func(t *testing.T) {
		result := addTwoNumbers(
			toList(5, 2, 3),
			toList(7, 8),
		)
		assertEqual(t, []int{2, 1, 4}, result)
	})

	t.Run("zero in middle", func(t *testing.T) {
		result := addTwoNumbers(
			toList(1, 6, 0, 3),
			toList(6, 3, 0, 8, 9),
		)
		assertEqual(t, []int{7, 9, 0, 1, 0, 1}, result)
	})
}
