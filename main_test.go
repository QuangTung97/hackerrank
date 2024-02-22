package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoman(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert.Equal(t, "III", intToRoman(3))
		assert.Equal(t, "LVIII", intToRoman(58))
		assert.Equal(t, "MCMXCIV", intToRoman(1994))
	})
}
