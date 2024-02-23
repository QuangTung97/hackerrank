package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
	assert.Equal(t, []string{
		"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
	}, letterCombinations("23"))
}
