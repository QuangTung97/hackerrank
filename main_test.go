package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeros(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert.Equal(t, 1, minDistance("a", ""))
	})

	t.Run("more complex", func(t *testing.T) {
		assert.Equal(t, 2, minDistance("aa", ""))
	})

	t.Run("more complex", func(t *testing.T) {
		assert.Equal(t, 1, minDistance("aa", "ab"))
	})

	t.Run("ex1", func(t *testing.T) {
		assert.Equal(t, 3, minDistance("horse", "ros"))
	})

	t.Run("ex2", func(t *testing.T) {
		assert.Equal(t, 5, minDistance("intention", "execution"))
	})
}
