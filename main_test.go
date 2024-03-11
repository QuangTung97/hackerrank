package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		p := position{
			row: 0,
			col: 0,
		}
		w0 := window{
			top:   -1,
			left:  -1,
			right: 3,
			bot:   2,
		}
		d0 := position{
			row: 0,
			col: 1,
		}

		p, d1, w1 := computeNext(p, d0, w0)
		assert.Equal(t, w0, w1)
		assert.Equal(t, d0, d1)
		assert.Equal(t, position{row: 0, col: 1}, p)

		p, d1, w1 = computeNext(p, d0, w0)
		assert.Equal(t, w0, w1)
		assert.Equal(t, d0, d1)
		assert.Equal(t, position{row: 0, col: 2}, p)

		p, d2, w2 := computeNext(p, d1, w0)
		w1.top++
		assert.Equal(t, w1, w2)
		assert.Equal(t, position{row: 1, col: 0}, d2)
		assert.Equal(t, position{row: 1, col: 2}, p)

		p, d3, w3 := computeNext(p, d2, w2)
		w2.right--
		assert.Equal(t, w2, w3)
		assert.Equal(t, position{row: 0, col: -1}, d3)
		assert.Equal(t, position{row: 1, col: 1}, p)

		p, d4, w4 := computeNext(p, d3, w3)
		assert.Equal(t, w3, w4)
		assert.Equal(t, position{row: 0, col: -1}, d4)
		assert.Equal(t, position{row: 1, col: 0}, p)
	})

	t.Run("square", func(t *testing.T) {
		p := position{
			row: 0,
			col: 0,
		}
		w0 := window{
			top:   -1,
			left:  -1,
			right: 3,
			bot:   3,
		}
		d0 := position{
			row: 0,
			col: 1,
		}

		p, d1, w1 := computeNext(p, d0, w0)
		assert.Equal(t, w0, w1)
		assert.Equal(t, d0, d1)
		assert.Equal(t, position{row: 0, col: 1}, p)

		p, d2, w2 := computeNext(p, d0, w0)
		assert.Equal(t, w0, w2)
		assert.Equal(t, d0, d2)
		assert.Equal(t, position{row: 0, col: 2}, p)

		p, d3, w3 := computeNext(p, d0, w2)
		w2.top++
		assert.Equal(t, w2, w3)
		assert.Equal(t, position{row: 1, col: 0}, d3)
		assert.Equal(t, position{row: 1, col: 2}, p)

		p, d4, w4 := computeNext(p, d3, w3)
		assert.Equal(t, w3, w4)
		assert.Equal(t, d3, d4)
		assert.Equal(t, position{row: 2, col: 2}, p)

		p, d5, w5 := computeNext(p, d4, w4)
		w4.right--
		assert.Equal(t, w4, w5)
		assert.Equal(t, position{row: 0, col: -1}, d5)
		assert.Equal(t, position{row: 2, col: 1}, p)

		p, d6, w6 := computeNext(p, d5, w5)
		assert.Equal(t, w5, w6)
		assert.Equal(t, position{row: 0, col: -1}, d6)
		assert.Equal(t, position{row: 2, col: 0}, p)

		p, d7, w7 := computeNext(p, d6, w6)
		w6.bot--
		assert.Equal(t, w6, w7)
		assert.Equal(t, position{row: -1, col: 0}, d7)
		assert.Equal(t, position{row: 1, col: 0}, p)

		p, d8, w8 := computeNext(p, d7, w7)
		w7.left++
		assert.Equal(t, w7, w8)
		assert.Equal(t, position{row: 0, col: 1}, d8)
		assert.Equal(t, position{row: 1, col: 1}, p)
	})
}

func TestSpiralOrder(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}, generateMatrix(3))
}
