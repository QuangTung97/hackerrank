package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	a := []int{3, 2, 2, 3}
	index := removeElement(a, 3)
	assert.Equal(t, []int{2, 2}, a[:index])
}
