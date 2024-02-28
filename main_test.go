package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {
	assert.Equal(t, "48", multiply("12", "4"))
}

func TestMul_Case2(t *testing.T) {
	assert.Equal(t, "52", multiply("13", "4"))
}

func TestMul_Case3(t *testing.T) {
	assert.Equal(t, "56088", multiply("123", "456"))
}

func TestMul_Case4(t *testing.T) {
	assert.Equal(t, "81", multiply("9", "9"))
}

func TestMul_Case5(t *testing.T) {
	assert.Equal(t, "998001", multiply("999", "999"))
}

func TestMul_Case6(t *testing.T) {
	assert.Equal(t, "9801", multiply("99", "99"))
}

func TestMul_Case7(t *testing.T) {
	assert.Equal(t, "0", multiply("9133", "0"))
}
