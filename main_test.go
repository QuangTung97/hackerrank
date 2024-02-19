package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindConcat(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		result := findConcatSubstrings([]string{"a", "b", "c", "b", "a", "d"}, []string{"a", "b"})
		assert.Equal(t, []int{0, 3}, result)
	})

	t.Run("case2", func(t *testing.T) {
		result := findConcatSubstrings([]string{"c", "a", "b"}, []string{"a", "b"})
		assert.Equal(t, []int{1}, result)
	})

	t.Run("case3", func(t *testing.T) {
		result := findConcatSubstrings([]string{"a", "b", "c", "b", "a", "d"}, []string{"a", "b", "c"})
		assert.Equal(t, []int{0, 2}, result)
	})
}

func TestFindSubstring(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		res := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
		assert.Equal(t, []int{0, 9}, res)
	})
	t.Run("normal", func(t *testing.T) {
		res := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
		assert.Equal(t, []int{}, res)
	})
}
