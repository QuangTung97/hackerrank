package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}))
	})
	t.Run("false", func(t *testing.T) {
		assert.Equal(t, false, canJump([]int{3, 2, 1, 0, 4}))
	})
	t.Run("false", func(t *testing.T) {
		assert.Equal(t, true, canJump([]int{0}))
	})
}
