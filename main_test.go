package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, true, isValid("[]"))
	assert.Equal(t, false, isValid("[)"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, true, isValid("{[]}"))
	assert.Equal(t, false, isValid("["))
	assert.Equal(t, false, isValid("]"))
}
