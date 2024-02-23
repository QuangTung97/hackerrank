package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
				continue
			}
			if sum > 0 {
				r--
				continue
			}
			result = append(result, []int{nums[i], nums[l], nums[r]})
			l++
			for l < r {
				if nums[l] != nums[l-1] {
					break
				}
				l++
			}
		}
	}
	return result
}
