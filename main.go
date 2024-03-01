package main

func lowerBound(intervals [][]int, x int) int {
	first := 0
	last := len(intervals)
	for first < last {
		mid := first + (last-first)/2
		if x <= intervals[mid][0] {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func combine(intervals [][]int) [][]int {
	var result [][]int
	for _, inv := range intervals {
		if len(result) > 0 {
			last := result[len(result)-1]
			if last[1] >= inv[0] {
				last[1] = max(last[1], inv[1])
				continue
			}
		}
		result = append(result, inv)
	}
	return result
}

func insert(input [][]int, newInterval []int) [][]int {
	index := lowerBound(input, newInterval[0])

	intervals := make([][]int, 0, len(input)+1)
	intervals = append(intervals, input[:index]...)
	intervals = append(intervals, newInterval)
	intervals = append(intervals, input[index:]...)

	return combine(intervals)
}
