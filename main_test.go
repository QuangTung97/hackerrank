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

func TestCombine(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 2}, {1, 3}, {1, 4},
		{2, 3}, {2, 4},
		{3, 4},
	}, combine(4, 2))
}

func TestCombine_Ex2(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
	}, combine(4, 3))
}
