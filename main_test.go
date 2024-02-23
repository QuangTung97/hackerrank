package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToList(t *testing.T) {
	h := newList(1, 2, 3, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, toInts(h))
}

func TestRotate(t *testing.T) {
	h := newList(1, 2, 3, 4, 5)
	newL := rotateRight(h, 2)
	assert.Equal(t, []int{4, 5, 1, 2, 3}, toInts(newL))
}

func TestRotate_Case2(t *testing.T) {
	h := newList(1)
	newL := rotateRight(h, 0)
	assert.Equal(t, []int{1}, toInts(newL))
}
