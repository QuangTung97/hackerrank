package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
	assert.Equal(t, "bab", longestPalindrome("babad"))
}
