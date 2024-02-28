package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := newList(1, 2, 3, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, toInts(l))
}

func TestSwap(t *testing.T) {
	l := swapPairs(newList(1, 2, 3, 4))
	assert.Equal(t, []int{2, 1, 4, 3}, toInts(l))
}

func TestSwap_2(t *testing.T) {
	l := swapPairs(newList(1, 2, 3, 4, 5))
	assert.Equal(t, []int{2, 1, 4, 3, 5}, toInts(l))
}
