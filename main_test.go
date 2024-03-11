package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeros(t *testing.T) {
	m := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(m)
	assert.Equal(t, [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}, m)
}

func TestSetZeros_Ex2(t *testing.T) {
	m := [][]int{
		{0},
		{1},
	}
	setZeroes(m)
	assert.Equal(t, [][]int{
		{0},
		{0},
	}, m)
}
