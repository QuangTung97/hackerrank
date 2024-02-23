package main

import (
	"slices"
)

func nextPermutation(nums []int) {
	mid := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			mid = i
			break
		}
	}

	if mid < 0 {
		slices.Reverse(nums)
		return
	}

	right := mid
	for i := mid + 1; i < len(nums); i++ {
		if nums[i] <= nums[mid-1] {
			break
		}
		right = i
	}
	nums[mid-1], nums[right] = nums[right], nums[mid-1]

	slices.Reverse(nums[mid:])
}
