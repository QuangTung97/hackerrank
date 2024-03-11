package main

func computeNextCounts(newCount []int, oldCount []int, obstacles []int) {
	for i := range newCount {
		newCount[i] = 0
		if i > 0 {
			newCount[i] = newCount[i-1]
		}
		newCount[i] += oldCount[i]
		if obstacles[i] > 0 {
			newCount[i] = 0
		}
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	oldCount := make([]int, len(obstacleGrid[0]))
	newCount := make([]int, len(obstacleGrid[0]))

	for i := range oldCount {
		if obstacleGrid[0][i] > 0 {
			break
		}
		oldCount[i] = 1
	}

	for row := 1; row < len(obstacleGrid); row++ {
		computeNextCounts(newCount, oldCount, obstacleGrid[row])
		tmp := oldCount
		oldCount = newCount
		newCount = tmp
	}

	n := len(obstacleGrid[0])
	return oldCount[n-1]
}
