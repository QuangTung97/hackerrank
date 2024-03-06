package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArr(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 1, maxSubArray([]int{1}))
}

func TestExample1(t *testing.T) {
	arr := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}
	assert.Equal(t, 6, maxSubArray(arr))
}
