package main

import (
	"slices"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	slices.Sort(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r] - target
				if sum < 0 {
					l++
					continue
				}
				if sum > 0 {
					r--
					continue
				}
				result = append(result, []int{
					nums[i], nums[j], nums[l], nums[r],
				})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return result
}
