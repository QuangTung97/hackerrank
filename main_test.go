package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitSet(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		var b bitset
		b.add(0)
		assert.Equal(t, [2]uint64{1, 0}, b.data)

		b.add(2)
		assert.Equal(t, [2]uint64{0x5, 0}, b.data)

		assert.Equal(t, false, b.existed(1))
		assert.Equal(t, true, b.existed(2))

		b.add(63)
		assert.Equal(t, [2]uint64{0x8000_0000_0000_0005, 0}, b.data)

		assert.Equal(t, false, b.existed(62))
		assert.Equal(t, true, b.existed(63))

		b.add(64)
		b.add(66)
		assert.Equal(t, [2]uint64{0x8000_0000_0000_0005, 0x5}, b.data)

		assert.Equal(t, true, b.existed(66))
		assert.Equal(t, false, b.existed(67))
	})
}

func TestLongestSubString(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
}

func TestLongestSubString_Start(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abc"))
}
