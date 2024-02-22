package main

import (
	"sort"
)

func lowerBound(nums []int, x int) int {
	first := 0
	last := len(nums)
	for first < last {
		mid := first + (last-first)/2
		if nums[mid] < x {
			first = mid + 1
		} else {
			last = mid
		}
	}
	return first
}

type triplet struct {
	a int
	b int
	c int
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	existed := map[triplet]struct{}{}
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			twoSum := nums[i] + nums[j]

			searchNums := nums[j+1:]
			for {
				index := lowerBound(searchNums, -twoSum)
				index2 := lowerBound(searchNums, -twoSum+1)
				if index >= len(searchNums) || searchNums[index] != -twoSum {
					break
				}

				newTuple := []int{nums[i], nums[j], searchNums[index]}
				t := triplet{
					a: newTuple[0],
					b: newTuple[1],
					c: newTuple[2],
				}
				_, ok := existed[t]
				if !ok {
					existed[t] = struct{}{}
					result = append(result, newTuple)
				}
				searchNums = searchNums[index2:]
			}
		}
	}
	return result
}
