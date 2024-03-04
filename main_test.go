package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPos(t *testing.T) {
	assert.Equal(t,
		position{rows: 0, cols: 2},
		nextPos(position{rows: 0, cols: 0}, 3),
	)

	assert.Equal(t,
		position{rows: 2, cols: 2},
		nextPos(position{rows: 0, cols: 2}, 3),
	)

	assert.Equal(t,
		position{rows: 2, cols: 0},
		nextPos(position{rows: 2, cols: 2}, 3),
	)
}

func TestNextPos_Bigger_Space(t *testing.T) {
	assert.Equal(t,
		position{rows: 2, cols: 2},
		nextPos(position{rows: 1, cols: 2}, 4),
	)
}

func TestSample(t *testing.T) {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(m)
	assert.Equal(t, [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}, m)
}

func TestSample_Bigger(t *testing.T) {
	m := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(m)
	assert.Equal(t, [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}, m)
}
