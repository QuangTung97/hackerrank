package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	t.Run("3", func(t *testing.T) {
		assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	})

	t.Run("4", func(t *testing.T) {
		assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	})

	t.Run("1", func(t *testing.T) {
		assert.Equal(t, "A", convert("A", 1))
	})

	t.Run("special", func(t *testing.T) {
		assert.Equal(t, "ABDC", convert("ABCD", 3))
	})
}
