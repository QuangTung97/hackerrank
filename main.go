package main

import (
	"math"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := math.MaxInt
	result := 0

	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum3 := nums[i] + nums[l] + nums[r]
			sum := sum3 - target

			if abs(sum) < closest {
				closest = abs(sum)
				result = sum3
			}
			if sum < 0 {
				l++
				continue
			}
			if sum > 0 {
				r--
				continue
			}
			return sum3
		}
	}

	return result
}
