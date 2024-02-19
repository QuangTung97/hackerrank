package main

func findMin(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] <= 0 {
			continue
		}
		index := arr[i] - 1
		arr[i] = 0
		for {
			if index >= len(arr) {
				break
			}
			nextNum := arr[index]
			arr[index] = -1
			if nextNum <= 0 {
				break
			}
			index = nextNum - 1
		}
	}
	for i := range arr {
		if arr[i] == 0 {
			return i + 1
		}
	}
	return len(arr) + 1
}

func firstMissingPositive(nums []int) int {
	last := len(nums)
	for i := 0; i < last; {
		if nums[i] <= 0 {
			last--
			nums[i] = nums[last]
			continue
		}
		i++
	}
	nums = nums[:last]
	return findMin(nums)
}
