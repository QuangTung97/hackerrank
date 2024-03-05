package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoverRange(t *testing.T) {
	assert.Equal(t, 5, coverRange(0, 5))
	assert.Equal(t, 5, coverRange(1, 5))
}

func TestBuildSegmentTree(t *testing.T) {
	nodes := buildTree([]int{5, 4, 8, 3})
	assert.Equal(t, []int{3, 4, 3, 5, 4, 8, 3}, nodes)

	assert.Equal(t, 4, treeMinOf(nodes, 0))
}

func TestJumpGame(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
}
