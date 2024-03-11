package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeros(t *testing.T) {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	assert.Equal(t, true, searchMatrix(m, 3))
	assert.Equal(t, false, searchMatrix(m, 4))
	assert.Equal(t, false, searchMatrix(m, 18))
}

func TestSetZeros_Ex1(t *testing.T) {
	m := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	assert.Equal(t, true, searchMatrix(m, 3))
}
