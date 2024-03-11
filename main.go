package main

func computeNext(newSum []int, oldSum []int, row []int) {
	for i := range newSum {
		if i > 0 {
			newSum[i] = min(newSum[i-1], oldSum[i])
		} else {
			newSum[i] = oldSum[i]
		}
		newSum[i] += row[i]
	}
}
func minPathSum(grid [][]int) int {
	oldSum := make([]int, len(grid[0]))
	newSum := make([]int, len(grid[0]))

	for i := range oldSum {
		if i == 0 {
			oldSum[i] = grid[0][0]
		} else {
			oldSum[i] = oldSum[i-1] + grid[0][i]
		}
	}

	for row := 1; row < len(grid); row++ {
		computeNext(newSum, oldSum, grid[row])
		tmp := oldSum
		oldSum = newSum
		newSum = tmp
	}

	return oldSum[len(grid[0])-1]
}
