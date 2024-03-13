package main

import (
	"slices"
)

func nextIter(it []int, n int) bool {
	for i := len(it) - 1; i >= 0; i-- {
		right := len(it) - 1 - i
		if it[i] >= n-right {
			continue
		}

		next := it[i]
		for ; i < len(it); i++ {
			next++
			it[i] = next
		}
		return true
	}
	return false
}

func combine(n int, k int) [][]int {
	var result [][]int

	iter := make([]int, k)
	for i := range iter {
		iter[i] = i + 1
	}

	for {
		result = append(result, slices.Clone(iter))
		ok := nextIter(iter, n)
		if !ok {
			break
		}
	}

	return result
}
