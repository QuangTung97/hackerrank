package main

func removeDuplicates(nums []int) int {
	left := 1
	duplicateCount := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[left-1] {
			duplicateCount++
			if duplicateCount >= 2 {
				duplicateCount = 1
				continue
			}
		} else {
			duplicateCount = 0
		}

		nums[left] = nums[i]
		left++
	}

	return left
}
