package main

import (
	"slices"
)

func intersect(a, b []int) bool {
	if b[0] <= a[1] {
		return true
	}
	return false
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	var result [][]int
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intersect(last, intervals[i]) {
			last[1] = max(last[1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
