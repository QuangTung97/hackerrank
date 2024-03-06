package main

func maxSubArray(nums []int) int {
	res := nums[0]
	for start := 0; start < len(nums); start++ {
		res = max(res, nums[start])
		if nums[start] <= 0 {
			continue
		}
		if start+1 >= len(nums) {
			continue
		}
		nums[start+1] += nums[start]
	}
	return res
}
