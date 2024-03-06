package main

func maxSubArray(nums []int) int {
	res := nums[0]
	prefixSum := 0
	for start := 0; start < len(nums); start++ {
		prefixSum += nums[start]
		res = max(res, prefixSum)
		if prefixSum <= 0 {
			prefixSum = 0
			continue
		}
	}
	return res
}
