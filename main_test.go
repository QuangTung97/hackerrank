package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}
		n := removeDuplicates(nums)
		assert.Equal(t, []int{1, 1, 2, 2, 3}, nums[:n])
	})

	t.Run("normal", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		n := removeDuplicates(nums)
		assert.Equal(t, []int{0, 0, 1, 1, 2, 3, 3}, nums[:n])
	})

	t.Run("ex1", func(t *testing.T) {
		nums := []int{1, 2, 2}
		n := removeDuplicates(nums)
		assert.Equal(t, []int{1, 2, 2}, nums[:n])
	})

	t.Run("ex2", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		n := removeDuplicates(nums)
		assert.Equal(t, []int{0, 0, 1, 1, 2, 3, 3}, nums[:n])
	})
}
