package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, 1, maxArea([]int{1, 1}))
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 4, maxArea([]int{1, 2, 4, 3}))
}

func TestExample_err(t *testing.T) {
	assert.Equal(t, 4, maxArea([]int{1, 2, 4, 3}))
}
