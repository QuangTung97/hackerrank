package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	stepSize := (numRows - 1) + (numRows - 2) + 1

	var buf strings.Builder
	for row := 0; row < numRows; row++ {
		for startIndex := 0; startIndex < len(s)+stepSize; startIndex += stepSize {
			if startIndex > 0 && row > 0 && row < numRows-1 {
				leftIndex := startIndex - row
				if leftIndex < len(s) {
					buf.WriteByte(s[leftIndex])
				}
			}
			i := startIndex + row
			if i < len(s) {
				buf.WriteByte(s[i])
			}
		}
	}
	return buf.String()
}
