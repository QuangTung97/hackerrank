package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower(t *testing.T) {
	assert.Equal(t, 4.0, myPow(2.0, 2))
	assert.Equal(t, 9.0, myPow(3.0, 2))
	assert.Equal(t, 27.0, myPow(3.0, 3))
	assert.Equal(t, 1.0/27.0, myPow(3.0, -3))
}
