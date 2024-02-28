package main

func countNumberOf(n int, x int) int {
	count := 0
	base := x
	for base <= n {
		count += n / base
		base *= x
	}
	return count
}

func trailingZeroes(n int) int {
	return min(countNumberOf(n, 2), countNumberOf(n, 5))
}
