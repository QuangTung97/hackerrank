package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi("  -42"))
	assert.Equal(t, 42, myAtoi("  +42"))
	assert.Equal(t, 4193, myAtoi("  4193 with words"))

	assert.Equal(t, 0, myAtoi(""))
	assert.Equal(t, 2147483647, myAtoi("18446744073709551617"))
}
