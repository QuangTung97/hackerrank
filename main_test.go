package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	assert.Equal(t, 0, countNumberOf(1, 2))
	assert.Equal(t, 8, countNumberOf(10, 2))
	assert.Equal(t, 97, countNumberOf(100, 2))
	assert.Equal(t, 25, countNumberOf(28, 2))
	assert.Equal(t, 127, countNumberOf(128, 2))
}

func TestTrailingZero(t *testing.T) {
	assert.Equal(t, 0, trailingZeroes(3))
	assert.Equal(t, 1, trailingZeroes(5))
	assert.Equal(t, 7, trailingZeroes(30))
}

func TestTrailingZero_1(t *testing.T) {
	assert.Equal(t, 7, trailingZeroes(30))
}
