package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 5},
		{6, 9},
	}, insert([][]int{
		{1, 3},
		{6, 9},
	}, []int{2, 5}))
}
