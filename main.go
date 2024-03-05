package main

import (
	"math"
)

func bitLen(a int) int {
	if a == 0 {
		return 1
	}

	n := 0
	for a > 0 {
		n++
		a = a >> 1
	}
	return n
}

func divide(dividend int, divisor int) int {
	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}

	aLen := bitLen(dividend)
	bLen := bitLen(divisor)
	if aLen < bLen {
		return 0
	}

	remain := aLen - bLen
	result := 0
	for shift := remain; shift >= 0; shift-- {
		result <<= 1

		v := divisor << shift
		if v > dividend {
			continue
		}
		dividend -= v
		result |= 1
	}
	if sign < 0 {
		result = -result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return result
}
