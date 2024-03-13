package main

import (
	"math/bits"
)

func subsets(nums []int) [][]int {
	n := len(nums)
	maxVal := uint16(1 << n)
	result := make([][]int, 0, maxVal)
	for i := uint16(0); i < maxVal; i++ {
		value := make([]int, 0, bits.OnesCount16(i))
		mask := uint16(1)
		for offset := uint16(0); offset < uint16(n); offset++ {
			if i&mask != 0 {
				value = append(value, nums[offset])
			}
			mask <<= 1
		}

		result = append(result, value)
	}
	return result
}
