package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubset(t *testing.T) {
	assert.Equal(t, [][]int{
		{},
		{2}, {3},
		{2, 3},
	}, subsets([]int{2, 3}))
}
