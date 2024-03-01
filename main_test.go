package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}, permute([]int{1, 2, 3}))
}

func TestNext(t *testing.T) {
	a := []int{1, 2, 3}

	assert.Equal(t, true, nextPerm(a))
	assert.Equal(t, []int{1, 3, 2}, a)

	assert.Equal(t, true, nextPerm(a))
	assert.Equal(t, []int{2, 1, 3}, a)

	assert.Equal(t, true, nextPerm(a))
	assert.Equal(t, []int{2, 3, 1}, a)

	assert.Equal(t, true, nextPerm(a))
	assert.Equal(t, []int{3, 1, 2}, a)
}
