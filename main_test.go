package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert.Equal(t, 2, findMin([]int{1, 3, 4}))
	})

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 1, findMin([]int{}))
	})

	t.Run("first", func(t *testing.T) {
		assert.Equal(t, 1, findMin([]int{2, 3, 4}))
	})

	t.Run("after last", func(t *testing.T) {
		assert.Equal(t, 5, findMin([]int{1, 2, 3, 4}))
	})
}

func TestEx(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, 3, firstMissingPositive([]int{1, 2, 0}))
	})
}
