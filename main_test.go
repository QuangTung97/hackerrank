package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func TestSearch_Case2(t *testing.T) {
	assert.Equal(t, [][]int{
		{2, 2, 2, 2},
	}, fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func TestSearch_Case3(t *testing.T) {
	assert.Equal(t, [][]int{
		{-2, -1, 1, 2},
		{-1, -1, 1, 1},
	}, fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}
