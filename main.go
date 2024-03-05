package main

type state struct {
	count [][]int
}

func newState(m int, n int) *state {
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	a[0][0] = 1
	return &state{
		count: a,
	}
}

func (s *state) uniquePaths(row, col int) int {
	if s.count[row][col] > 0 {
		return s.count[row][col]
	}

	top := 0
	if row > 0 {
		top = s.uniquePaths(row-1, col)
	}

	left := 0
	if col > 0 {
		left = s.uniquePaths(row, col-1)
	}

	res := top + left
	s.count[row][col] = res
	return res
}

func uniquePaths(m int, n int) int {
	s := newState(m, n)
	res := s.uniquePaths(m-1, n-1)
	return res
}
