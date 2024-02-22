package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, 1, lowerBound([]int{1, 2, 3, 5, 6, 6, 7}, 2))
	assert.Equal(t, 6, lowerBound([]int{1, 2, 3, 5, 6, 6, 7}, 7))
	assert.Equal(t, 4, lowerBound([]int{1, 2, 3, 5, 6, 6, 7}, 6))
	assert.Equal(t, 0, lowerBound([]int{1, 1, 1}, 1))
	assert.Equal(t, 3, lowerBound([]int{1, 1, 1}, 3))
}
func Test3Sum(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	})

	t.Run("simple", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{0, 0, 0},
		}, threeSum([]int{0, 0, 0}))
	})

	t.Run("duplicated", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2},
		}, threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
	})
}
