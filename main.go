package main

func removeDuplicates(nums []int) int {
	nextIndex := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			nums[nextIndex] = nums[i]
			nextIndex++
		}
	}
	return nextIndex
}
