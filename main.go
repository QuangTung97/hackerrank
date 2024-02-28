package main

import (
	"strings"
)

func reverse(s string) string {
	var buf strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func multiply(num1 string, num2 string) string {
	num1 = reverse(num1)
	num2 = reverse(num2)

	result := make([]byte, 10000)
	last := 0

	rem := uint8(0)
	for i := 0; i < len(num2); i++ {
		b := num2[i] - '0'
		for j := 0; j < len(num1); j++ {
			a := num1[j] - '0'
			m := a*b + rem + result[i+j]
			rem = m / 10
			result[i+j] = m % 10
			if m > 0 {
				last = max(last, i+j+1)
			}
		}
		if rem > 0 {
			result[last] = rem
			rem = 0
			last++
		}
	}

	for i := range result[:last] {
		result[i] = result[i] + '0'
	}

	if last == 0 {
		return "0"
	}

	return reverse(string(result[:last]))
}
