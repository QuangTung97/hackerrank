package main

func canJump(nums []int) bool {
	isPossible := make([]bool, len(nums))
	isPossible[len(nums)-1] = true
Outer:
	for i := len(nums) - 1; i >= 0; i-- {
		maxIndex := i + nums[i]
		if maxIndex >= len(nums) {
			maxIndex = len(nums) - 1
		}
		for j := maxIndex; j > i; j-- {
			if isPossible[j] {
				isPossible[i] = true
				continue Outer
			}
		}
	}
	return isPossible[0]
}
