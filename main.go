package main

import (
	"slices"
)

type state struct {
}

func newState() *state {
	return &state{}
}

func (s *state) combineSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	var result [][]int
	for index, e := range candidates {
		if e > target {
			break
		}

		if index > 0 && candidates[index-1] == e {
			continue
		}

		if e == target {
			result = append(result, []int{
				target,
			})
		}

		newTarget := target - e
		if newTarget < e {
			continue
		}

		subList := s.combineSum(candidates[index+1:], target-e)
		for _, l := range subList {
			newList := make([]int, 0, len(l)+1)
			newList = append(newList, e)
			newList = append(newList, l...)
			result = append(result, newList)
		}
	}

	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	s := newState()
	return s.combineSum(candidates, target)
}
