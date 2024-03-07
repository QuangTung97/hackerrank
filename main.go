package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	type position struct {
		key   string
		index int
	}

	posList := make([]position, 0, len(strs))
	for index, s := range strs {
		key := []byte(s)
		slices.Sort(key)
		posList = append(posList, position{
			key:   string(key),
			index: index,
		})
	}

	slices.SortFunc(posList, func(a, b position) int {
		if a.key < b.key {
			return -1
		}
		if a.key == b.key {
			return 0
		}
		return 1
	})

	result := make([][]string, 0, len(strs)/2)
	result = append(result, []string{
		strs[posList[0].index],
	})

	lastKey := posList[0].key
	for _, pos := range posList[1:] {
		if pos.key == lastKey {
			n := len(result)
			result[n-1] = append(result[n-1], strs[pos.index])
			continue
		}
		lastKey = pos.key
		result = append(result, []string{
			strs[pos.index],
		})
	}

	return result
}
