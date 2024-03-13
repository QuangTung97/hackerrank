package main

func nextIter(it []int, n int) bool {
	for i := len(it) - 1; i >= 0; i-- {
		right := len(it) - 1 - i
		if it[i] >= n-right {
			continue
		}

		next := it[i]
		for ; i < len(it); i++ {
			next++
			it[i] = next
		}
		return true
	}
	return false
}

func subsets(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0, 1<<n)
	result = append(result, []int{})

	for size := 1; size <= n; size++ {
		iter := make([]int, size)
		for i := range iter {
			iter[i] = i
		}
		for {
			value := make([]int, 0, size)
			for i := 0; i < size; i++ {
				value = append(value, nums[iter[i]])
			}
			result = append(result, value)
			ok := nextIter(iter, n-1)
			if !ok {
				break
			}
		}
	}

	return result
}
