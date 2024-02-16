package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	arr := []int{1, 1, 2, 2, 2, 3, 4, 5, 5, 6}
	index := removeDuplicates(arr)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, arr[:index])
}
