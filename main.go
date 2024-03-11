package main

func searchRow(matrix [][]int, target int) int {
	n := len(matrix[0])
	first := 0
	last := len(matrix)
	for first < last {
		mid := first + (last-first)/2
		if target <= matrix[mid][n-1] {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func searchCol(row []int, target int) int {
	first := 0
	last := len(row)
	for first < last {
		mid := first + (last-first)/2
		if target <= row[mid] {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func searchMatrix(matrix [][]int, target int) bool {
	foundRow := searchRow(matrix, target)
	if foundRow >= len(matrix) {
		return false
	}

	index := searchCol(matrix[foundRow], target)
	if index >= len(matrix[foundRow]) {
		return false
	}
	return matrix[foundRow][index] == target
}
