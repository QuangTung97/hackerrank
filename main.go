package main

import (
	"fmt"
	"slices"
	"sort"
)

var visited = make([]bool, 50)

func findSmallerOrEq(nums []int, x int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= x {
			return i
		}
	}
	return -1
}

func computeRecur(candidates []int, target int) [][]int {
	fmt.Println(candidates, target)

	if len(candidates) == 0 || target == 0 {
		return [][]int{nil}
	}

	if visited[target] {
		return nil
	}
	visited[target] = true

	index := findSmallerOrEq(candidates, target)
	if index < 0 {
		return nil
	}

	var result [][]int
	for i := len(candidates) - 1; i >= 0; i-- {
		last := candidates[i]
		tmp := computeRecur(candidates[:i], target-last)

		newResult := make([][]int, 0, len(tmp))
		for _, e := range tmp {
			e = slices.Clone(e)
			newResult = append(
				newResult,
				append(e, last),
			)
		}
		result = append(result, newResult...)
	}
	fmt.Println(result, candidates, target)
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return computeRecur(candidates, target)
}
