package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	})
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, -2, threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
	})
}
