package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	nums := []int{1, 2, 3}

	nextPermutation(nums)
	assert.Equal(t, []int{1, 3, 2}, nums)

	nextPermutation(nums)
	assert.Equal(t, []int{2, 1, 3}, nums)

	nextPermutation(nums)
	assert.Equal(t, []int{2, 3, 1}, nums)

	nextPermutation(nums)
	assert.Equal(t, []int{3, 1, 2}, nums)

	nextPermutation(nums)
	assert.Equal(t, []int{3, 2, 1}, nums)

	nextPermutation(nums)
	assert.Equal(t, []int{1, 2, 3}, nums)
}

func TestFind_Special(t *testing.T) {
	nums := []int{1, 5, 1}
	nextPermutation(nums)
	assert.Equal(t, []int{5, 1, 1}, nums)
}
