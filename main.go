package main

import (
	"math"
)

func clearRow(matrix [][]int, row int, rootCol int) {
	for col := 0; col < len(matrix[0]); col++ {
		if col == rootCol {
			continue
		}
		if col > rootCol && matrix[row][col] == math.MaxInt {
			break
		}
		matrix[row][col] = 0
	}
}

func clearCol(matrix [][]int, col int, rootRow int) {
	for row := 0; row < len(matrix); row++ {
		if row == rootRow {
			continue
		}
		if row > rootRow && matrix[row][col] == math.MaxInt {
			break
		}
		matrix[row][col] = 0
	}
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				matrix[row][col] = math.MaxInt
			}
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == math.MaxInt {
				clearRow(matrix, row, col)
				clearCol(matrix, col, row)
			}
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == math.MaxInt {
				matrix[row][col] = 0
			}
		}
	}
}
