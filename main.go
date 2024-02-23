package main

func findRotatePoint(nums []int) int {
	first := 0
	last := len(nums) - 1
	for first < last-1 {
		mid := first + (last-first)/2
		if nums[first] > nums[mid] {
			last = mid
		} else {
			first = mid
		}
	}
	if nums[first] > nums[last] {
		return last
	}
	return first
}

func lowerBound(nums []int, a int) int {
	first := 0
	last := len(nums)
	for first < last {
		mid := first + (last-first)/2
		if a <= nums[mid] {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func search(nums []int, target int) int {
	rotate := findRotatePoint(nums)

	a := nums[:rotate]
	index := lowerBound(a, target)
	if index < len(a) && a[index] == target {
		return index
	}

	b := nums[rotate:]
	index = lowerBound(b, target)
	if index < len(b) && b[index] == target {
		return rotate + index
	}

	return -1
}
