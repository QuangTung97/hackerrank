package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSteps(t *testing.T) {
	assert.Equal(t, []step{
		{ch: 'a'},
	}, toSteps("a"))

	assert.Equal(t, []step{
		{ch: 'a', repeated: true},
	}, toSteps("a*"))

	assert.Equal(t, []step{
		{ch: 'a', repeated: false},
		{ch: 'b', repeated: true},
	}, toSteps("ab*"))

	assert.Equal(t, []step{
		{isAll: true, repeated: false},
	}, toSteps("."))
}

func TestAccept(t *testing.T) {
	t.Run("new state", func(t *testing.T) {
		steps := []step{
			{ch: 'a', repeated: true},
			{ch: 'a', repeated: false},
		}
		s := newState(steps)
		assert.Equal(t, &state{
			steps: steps,
			set: map[int]struct{}{
				0: {},
				1: {},
			},
		}, s)
	})

	t.Run("accept", func(t *testing.T) {
		steps := []step{
			{ch: 'a', repeated: true},
			{ch: 'a', repeated: false},
		}

		s := newState(steps)
		assert.Equal(t, map[int]struct{}{
			0: {}, 1: {},
		}, s.set)

		assert.Equal(t, true, s.accept('a'))
		assert.Equal(t, map[int]struct{}{
			0: {}, 1: {}, 2: {},
		}, s.set)

		assert.Equal(t, true, s.accept('a'))
		assert.Equal(t, map[int]struct{}{
			0: {}, 1: {}, 2: {},
		}, s.set)

		assert.Equal(t, false, s.accept('b'))
		assert.Equal(t, map[int]struct{}{}, s.set)

		assert.Equal(t, false, s.accept('b'))
		assert.Equal(t, map[int]struct{}{}, s.set)
	})
}

func TestExample(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert.Equal(t, false, isMatch("aa", "a"))
	})
	t.Run("star", func(t *testing.T) {
		assert.Equal(t, true, isMatch("aa", "a*"))
	})
	t.Run("dot", func(t *testing.T) {
		assert.Equal(t, true, isMatch("ab", ".*"))
	})
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, true, isMatch("aab", "c*a*b"))
	})
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, []step{
			{ch: 'a', repeated: true},
			{ch: 'a', repeated: false},
		}, toSteps("a*a"))
		assert.Equal(t, true, isMatch("aaa", "a*a"))
	})
	t.Run("not finish", func(t *testing.T) {
		assert.Equal(t, false, isMatch("ab", "abc"))
	})
}
