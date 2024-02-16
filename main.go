package main

import (
	"slices"
)

type verticalLine struct {
	height int
	index  int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func absInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func maxArea(height []int) int {
	ignored := make([]bool, len(height))

	lines := make([]verticalLine, 0, len(height))
	for i, h := range height {
		lines = append(lines, verticalLine{
			height: h,
			index:  i,
		})
	}
	slices.SortFunc(lines, func(a, b verticalLine) int {
		return b.height - a.height
	})

	maxVal := 0
	for i := 1; i < len(lines); i++ {
		a := lines[i]
		if ignored[a.index] {
			continue
		}
		for j := 0; j < i; j++ {
			b := lines[j]
			if ignored[b.index] {
				continue
			}

			minIndex := minInt(a.index, b.index)
			maxIndex := maxInt(a.index, b.index)

			for k := minIndex + 1; k < maxIndex; k++ {
				ignored[k] = true
			}

			area := minInt(a.height, b.height) * absInt(a.index, b.index)
			if maxVal < area {
				maxVal = area
			}
		}
	}
	return maxVal
}
