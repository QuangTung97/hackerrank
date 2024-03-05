package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
}

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, 1, uniquePaths(1, 1))
	assert.Equal(t, 28, uniquePaths(3, 7))
}
