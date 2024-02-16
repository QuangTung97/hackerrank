package main

import (
	"math"
)

func findFirstNonSpace(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			return i
		}
	}
	return 0
}

func isOutside(result uint64, isNeg bool) (int, bool) {
	if isNeg {
		if result > -math.MinInt32 {
			return math.MinInt32, true
		}
		return -int(result), false
	}
	if result > math.MaxInt32 {
		return math.MaxInt32, true
	}
	return int(result), false
}

func myAtoi(s string) int {
	start := findFirstNonSpace(s)

	if start >= len(s) {
		return 0
	}

	isNeg := false
	if s[start] == '-' {
		isNeg = true
		start++
	} else if s[start] == '+' {
		start++
	}

	result := uint64(0)
	for i := start; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			break
		}
		if i > 0 {
			result *= 10
		}
		ch := s[i] - '0'
		result += uint64(ch)
		if newVal, ok := isOutside(result, isNeg); ok {
			return newVal
		}
	}

	newVal, _ := isOutside(result, isNeg)
	return newVal
}
