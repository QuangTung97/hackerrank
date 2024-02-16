package main

func isValid(s string) bool {
	stack := make([]byte, 0, 1024)

	for i := 0; i < len(s); i++ {
		e := s[i]

		if e == '(' || e == '[' || e == '{' {
			stack = append(stack, e)
			continue
		}

		if e == ')' || e == ']' || e == '}' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if !isMatch(last, e) {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	return len(stack) == 0
}

func isMatch(a, b byte) bool {
	if a == '(' && b == ')' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	return false
}
