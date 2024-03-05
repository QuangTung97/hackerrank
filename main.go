package main

func removeElement(nums []int, val int) int {
	last := len(nums)
	for i := 0; i < last; {
		if nums[i] != val {
			i++
			continue
		}

		last--
		nums[i] = nums[last]
	}
	return last
}
