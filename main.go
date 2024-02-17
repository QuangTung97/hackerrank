package main

import (
	"sort"
)

func medianOf(nums []int) float64 {
	n := len(nums)
	index := n / 2
	if n%2 == 0 {
		return float64(nums[index]+nums[index-1]) / 2.0
	}
	return float64(nums[index])
}

func simpleForm(nums, simpler []int) float64 {
	if len(simpler) == 0 {
		return medianOf(nums)
	}

	x := simpler[0]

	if len(nums) == 0 {
		return float64(x)
	}

	n := len(nums)
	index := n / 2
	if n%2 == 0 {
		v1 := nums[index-1]
		v2 := nums[index]
		if x < v1 {
			return float64(v1)
		}
		if x > v2 {
			return float64(v2)
		}
		return float64(x)
	} else {
		v := nums[index]
		if len(nums) == 1 {
			return float64(v+x) / 2.0
		}

		v1 := nums[index-1]
		v2 := nums[index+1]

		if x < v1 {
			return float64(v1+v) / 2.0
		}
		if x > v2 {
			return float64(v+v2) / 2.0
		}

		return float64(v+x) / 2.0
	}
}

func form2Solution(nums1 []int, nums2 []int) float64 {
	var x []int
	x = append(x, nums1...)
	x = append(x, nums2...)
	sort.Ints(x)
	return medianOf(x)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) <= 1 {
		return simpleForm(nums2, nums1)
	}
	if len(nums2) <= 1 {
		return simpleForm(nums1, nums2)
	}

	if len(nums1) == 2 {
		return form2Solution(nums2, nums1)
	}
	if len(nums2) == 2 {
		return form2Solution(nums1, nums2)
	}

	lowerLen := func(x []int) int {
		return (len(x) - 1) / 2
	}

	smallerLen := min(lowerLen(nums1), lowerLen(nums2))
	m1 := medianOf(nums1)
	m2 := medianOf(nums2)

	if m1 < m2 {
		return findMedianSortedArrays(
			nums1[smallerLen:],
			nums2[:len(nums2)-smallerLen],
		)
	}
	return findMedianSortedArrays(
		nums1[:len(nums1)-smallerLen],
		nums2[smallerLen:],
	)
}
