package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotatePoint(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, 4, findRotatePoint([]int{4, 5, 6, 7, 0, 1, 2}))
		assert.Equal(t, 3, findRotatePoint([]int{4, 5, 6, 0, 1, 2}))
		assert.Equal(t, 1, findRotatePoint([]int{6, 0, 1, 2}))
	})
	t.Run("special", func(t *testing.T) {
		assert.Equal(t, 2, findRotatePoint([]int{9, 10, 0, 1, 2}))
	})
}

func TestFind(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
