package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	var maxVal = 0

	for left < right {
		current := min(height[left], height[right]) * (right - left)
		if current > maxVal {
			maxVal = current
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxVal
}
