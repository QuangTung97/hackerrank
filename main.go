package main

type step struct {
	ch       byte
	isAll    bool
	repeated bool
}

type state struct {
	steps []step
	set   map[int]struct{}
}

func (s *state) extendEpsilon() {
	for {
		var newElems []int
		for index := range s.set {
			if index >= len(s.steps) {
				continue
			}
			if s.steps[index].repeated {
				newElems = append(newElems, index+1)
			}
		}

		oldLen := len(s.set)
		for _, e := range newElems {
			s.set[e] = struct{}{}
		}
		if len(s.set) == oldLen {
			return
		}
	}
}

func newState(steps []step) *state {
	s := &state{
		steps: steps,
		set: map[int]struct{}{
			0: {},
		},
	}
	s.extendEpsilon()
	return s
}

func (s *state) accept(ch byte) bool {
	newSet := map[int]struct{}{}
	for index := range s.set {
		if index >= len(s.steps) {
			continue
		}

		st := s.steps[index]

		if st.isAll {
			if st.repeated {
				newSet[index] = struct{}{}
				continue
			}
			newSet[index+1] = struct{}{}
			continue
		}

		if st.ch == ch {
			if st.repeated {
				newSet[index] = struct{}{}
				continue
			}
			newSet[index+1] = struct{}{}
			continue
		}
	}

	s.set = newSet
	if len(s.set) == 0 {
		return false
	}

	s.extendEpsilon()
	return true
}

func toSteps(p string) []step {
	var result []step
	for i := 0; i < len(p); i++ {
		ch := p[i]

		if ch == '*' {
			result[len(result)-1].repeated = true
			continue
		}

		if ch == '.' {
			result = append(result, step{
				isAll: true,
			})
			continue
		}

		result = append(result, step{
			ch: ch,
		})
	}
	return result
}

func isMatch(s string, p string) bool {
	steps := toSteps(p)
	st := newState(steps)

	for i := range s {
		ch := s[i]
		if !st.accept(ch) {
			return false
		}
	}

	_, ok := st.set[len(steps)]
	return ok
}
