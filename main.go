package main

import (
	"slices"
)

type state struct {
	visited []bool
	output  [][][]int
	existed map[string]struct{}
}

func newState() *state {
	return &state{
		visited: make([]bool, 41),
		output:  make([][][]int, 41),
		existed: map[string]struct{}{},
	}
}

func numsToKey(nums []int) string {
	res := make([]byte, 0, len(nums))
	for _, n := range nums {
		res = append(res, byte(n)+'A')
	}
	slices.Sort(res)
	return string(res)
}

func (s *state) compute(candidates []int, target int) [][]int {
	if s.visited[target] {
		return s.output[target]
	}

	s.visited[target] = true

	var l [][]int

	for _, e := range candidates {
		if e > target {
			break
		}

		if target == e {
			l = append(l, []int{
				e,
			})
			continue
		}

		newTarget := target - e
		subLists := s.compute(candidates, newTarget)
		for _, subList := range subLists {
			res := make([]int, 0, len(subList)+1)
			res = append(res, e)
			res = append(res, subList...)

			key := numsToKey(res)
			_, ok := s.existed[key]
			if ok {
				continue
			}
			s.existed[key] = struct{}{}
			l = append(l, res)
		}
	}

	s.output[target] = l
	return s.output[target]
}

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	s := newState()
	return s.compute(candidates, target)
}
