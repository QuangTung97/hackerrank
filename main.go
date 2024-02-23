package main

func lowerBound(nums []int, x int) int {
	first := 0
	last := len(nums)
	for first < last {
		mid := first + (last-first)/2
		if x <= nums[mid] {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func upperBound(nums []int, x int) int {
	first := 0
	last := len(nums)
	for first < last {
		mid := first + (last-first)/2
		if x < nums[mid] {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	end := upperBound(nums, target)
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	return []int{start, end - 1}
}
