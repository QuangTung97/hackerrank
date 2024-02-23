package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		assert.Equal(t, []string{
			"((()))", "(()())", "(())()", "()(())", "()()()",
		}, generateParenthesis(3))
	})
}
