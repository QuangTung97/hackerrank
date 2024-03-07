package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArr(t *testing.T) {
	assert.Equal(t, [][]string{
		{"bat"},
		{"ate", "eat", "tea"},
		{"nat", "tan"},
	}, groupAnagrams([]string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}))
}
