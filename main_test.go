package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, []token{
		{name: "home"},
	}, convertToTokens("/home"))

	assert.Equal(t, []token{
		{name: "home"},
	}, convertToTokens("/home/"))

	assert.Equal(t, []token{
		{name: "home"},
		{name: "user01"},
	}, convertToTokens("/home/user01"))

	assert.Equal(t, []token{
		{name: "home"},
		{name: "user01"},
	}, convertToTokens("/home//user01"))

	assert.Equal(t, []token{
		{name: "home"},
		{name: ".."},
		{name: "user01"},
	}, convertToTokens("/home/../user01"))

	assert.Equal(t, []token{
		{name: "home"},
		{name: "."},
		{name: "user01"},
	}, convertToTokens("/home/./user01"))
}

func TestSimplify(t *testing.T) {
	p := simplifyPath("/home/tung/")
	assert.Equal(t, "/home/tung", p)

	p = simplifyPath("/home/./tung")
	assert.Equal(t, "/home/tung", p)

	p = simplifyPath("/home/user01/../tung/")
	assert.Equal(t, "/home/tung", p)

	p = simplifyPath("/..")
	assert.Equal(t, "/", p)
}
