package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArr(t *testing.T) {
	assert.Equal(t, [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}, combinationSum([]int{2, 3, 5}, 8))
}
