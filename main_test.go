package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineSum2_Simple(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{1, 1},
			{2},
		}, combinationSum2([]int{1, 1, 2, 3}, 2))
	})
}

func TestCombineSum2(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func TestCombineSum2_Ex2(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 1, 1, 5},
		{1, 1, 3, 3},
		{3, 5},
	}, combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8))
}

func TestCombineSum2_Ex3(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 1, 2, 2},
		{1, 1, 4},
		{1, 2, 3},
		{2, 2, 2},
		{2, 4},
	}, combinationSum2([]int{4, 4, 2, 1, 4, 2, 2, 1, 3}, 6))
}
