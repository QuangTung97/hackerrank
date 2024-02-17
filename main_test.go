package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeanOf(t *testing.T) {
	assert.Equal(t, 2.0, medianOf([]int{1, 2, 3}))
	assert.Equal(t, 2.5, medianOf([]int{1, 2, 3, 4}))
	assert.Equal(t, 1.0, medianOf([]int{1}))
	assert.Equal(t, 4.5, medianOf([]int{4, 5}))
}

func TestExample(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{1, 3},
			[]int{2},
		)
		assert.Equal(t, 2.0, v)
	})

	t.Run("case2", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{1, 2},
			[]int{3, 4},
		)
		assert.Equal(t, 2.5, v)
	})

	t.Run("case3", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{2},
			[]int{},
		)
		assert.Equal(t, 2.0, v)
	})

	t.Run("case4", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{3},
			[]int{-2, -1},
		)
		assert.Equal(t, -1.0, v)
	})

	t.Run("case5", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{2, 2, 4, 4},
			[]int{2, 2, 4, 4},
		)
		assert.Equal(t, 3.0, v)
	})

	t.Run("case6", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{1, 2},
			[]int{-1, 3},
		)
		assert.Equal(t, 1.5, v)
	})

	t.Run("case7", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{1},
			[]int{2, 3, 4},
		)
		assert.Equal(t, 2.5, v)
	})

	t.Run("case8", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{1, 2, 6, 7},
			[]int{3, 4, 5, 8},
		)
		assert.Equal(t, 4.5, v)
	})

	t.Run("case9", func(t *testing.T) {
		v := findMedianSortedArrays(
			[]int{0, 0, 0, 0, 0},
			[]int{-1, 0, 0, 0, 0, 0, 1},
		)
		assert.Equal(t, 0.0, v)
	})
}
