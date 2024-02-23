package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	h := newList(1, 2, 3, 4, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, toInts(h))
}

func TestRemove(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		h := newList(1, 2, 3, 4, 5)
		newL := removeNthFromEnd(h, 2)
		assert.Equal(t, []int{1, 2, 3, 5}, toInts(newL))
	})
}
