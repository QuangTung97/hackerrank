package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextIter(t *testing.T) {
	iter := []int{1, 2, 3}

	nextIter(iter, 4)
	assert.Equal(t, []int{1, 2, 4}, iter)

	nextIter(iter, 4)
	assert.Equal(t, []int{1, 3, 4}, iter)

	nextIter(iter, 4)
	assert.Equal(t, []int{2, 3, 4}, iter)
}

func TestNextEx2(t *testing.T) {
	var ok bool
	iter := []int{1, 2}

	nextIter(iter, 4)
	assert.Equal(t, []int{1, 3}, iter)

	nextIter(iter, 4)
	assert.Equal(t, []int{1, 4}, iter)

	ok = nextIter(iter, 4)
	assert.Equal(t, []int{2, 3}, iter)
	assert.Equal(t, true, ok)

	nextIter(iter, 4)
	assert.Equal(t, []int{2, 4}, iter)

	ok = nextIter(iter, 4)
	assert.Equal(t, []int{3, 4}, iter)
	assert.Equal(t, true, ok)

	ok = nextIter(iter, 4)
	assert.Equal(t, []int{3, 4}, iter)
	assert.Equal(t, false, ok)
}

func TestSubset(t *testing.T) {
	assert.Equal(t, [][]int{
		{},
		{1}, {2},
		{1, 2},
	}, subsets([]int{1, 2}))
}
