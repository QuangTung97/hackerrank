package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitLen(t *testing.T) {
	assert.Equal(t, 1, bitLen(0))
	assert.Equal(t, 1, bitLen(1))
	assert.Equal(t, 2, bitLen(2))
	assert.Equal(t, 4, bitLen(15))
	assert.Equal(t, 5, bitLen(16))
}

func TestRemove(t *testing.T) {
	assert.Equal(t, 3, divide(10, 3))
	assert.Equal(t, -2, divide(7, -3))
	assert.Equal(t, 3, divide(25, 8))
	assert.Equal(t, 4, divide(32, 8))
	assert.Equal(t, 3, divide(31, 8))
	assert.Equal(t, 10, divide(31, 3))
	assert.Equal(t, 10, divide(32, 3))
	assert.Equal(t, 11, divide(33, 3))
	assert.Equal(t, -10, divide(-32, 3))
	assert.Equal(t, -11, divide(-33, 3))
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
}
