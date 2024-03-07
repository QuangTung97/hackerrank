package main

import (
	"slices"
)

func strToBits(s string) [52]uint16 {
	var res [52]uint16
	for i := 0; i < len(s); i++ {
		ch := s[i]
		pos := ch - 'A'
		if ch >= 'a' {
			pos = ch - 'a' + 26
		}
		res[pos]++
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	type position struct {
		key   [52]uint16
		index int
	}

	posList := make([]position, 0, len(strs))
	for index, s := range strs {
		posList = append(posList, position{
			key:   strToBits(s),
			index: index,
		})
	}

	slices.SortFunc(posList, func(a, b position) int {
		return slices.Compare(a.key[:], b.key[:])
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
