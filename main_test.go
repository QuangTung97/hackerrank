package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, [][]int{
		{2, 2, 3},
		{7},
	}, combinationSum([]int{2, 3, 6, 7}, 7))
}
