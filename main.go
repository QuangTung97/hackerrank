package main

func minDistance(word1 string, word2 string) int {
	oldTable := make([]int, len(word1)+1)
	newTable := make([]int, len(word1)+1)
	for i := range oldTable {
		oldTable[i] = i
	}

	for j := 1; j <= len(word2); j++ {
		for i := 0; i <= len(word1); i++ {
			if i == 0 {
				newTable[i] = j
				continue
			}
			ch1 := word1[i-1]
			ch2 := word2[j-1]
			if ch1 == ch2 {
				newTable[i] = oldTable[i-1]
				continue
			}

			newTable[i] = min(newTable[i-1], oldTable[i-1], oldTable[i]) + 1
		}
		tmp := oldTable
		oldTable = newTable
		newTable = tmp
	}
	return oldTable[len(word1)]
}
