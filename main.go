package main

import (
	"slices"
)

func nextPerm(nums []int) bool {
	findIndex := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			findIndex = i
			break
		}
	}
	if findIndex < 0 {
		return false
	}

	swapIndex := len(nums) - 1
	for j := findIndex + 2; j < len(nums); j++ {
		if nums[j] <= nums[findIndex] {
			swapIndex = j - 1
			break
		}
	}

	nums[findIndex], nums[swapIndex] = nums[swapIndex], nums[findIndex]
	slices.Reverse(nums[findIndex+1:])
	return true
}

func permute(nums []int) [][]int {
	slices.Sort(nums)
	var result [][]int
	for {
		result = append(result, slices.Clone(nums))
		if !nextPerm(nums) {
			break
		}
	}
	return result
}
