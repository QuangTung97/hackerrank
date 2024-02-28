package main

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		sum := numbers[l] + numbers[r] - target
		if sum < 0 {
			l++
			continue
		}
		if sum > 0 {
			r--
			continue
		}
		return []int{l + 1, r + 1}
	}
	return nil
}
