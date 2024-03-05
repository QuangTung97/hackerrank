package main

import (
	"math"
)

func computeInit(nums []int) []int {
	n := len(nums)
	size := 0
	for n > 1 {
		size += n
		n = (n + 1) / 2
	}
	size++
	nodes := make([]int, size)

	copy(nodes[size-len(nums):], nums)
	return nodes
}

func coverRange(i int, n int) int {
	return 0
}

func buildTree(nums []int) []int {
	nodes := computeInit(nums)

	n := len(nums)
	start := len(nodes) - n
	for {
		n = (n + 1) / 2
		oldStart := start
		start -= n

		for i := 0; i < n; i++ {
			l := oldStart + 2*i
			r := oldStart + 2*i + 1
			nodes[start+i] = min(nodes[l], nodes[r])
		}
		if start == 0 {
			break
		}
	}

	return nodes
}

func treeMinOf(nodes []int, end int) int {
	return 0
}

func jump(nums []int) int {
	minArr := make([]int, len(nums))
	for i := range minArr {
		minArr[i] = math.MaxInt32
	}
	minArr[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			step := nums[j]
			if j+step >= i {
				minArr[i] = min(minArr[i], minArr[j]+1)
			}
		}
	}
	return minArr[len(minArr)-1]
}
