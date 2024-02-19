package main

func findConcatSubstrings(words []string, match []string) []int {
	conf := map[string]int{}
	for _, key := range match {
		prev := conf[key]
		conf[key] = prev + 1
	}

	var result []int

	counter := map[string]int{}
	start := 0
	for end := 0; end < len(words); end++ {
		// fmt.Println(counter, start, end, result)
		w := words[end]
		cmp := conf[w]

		counter[w] = counter[w] + 1
		newCount := counter[w]

		if newCount > cmp {
			for ; ; start++ {
				startW := words[start]
				counter[startW] = counter[startW] - 1
				if startW == w {
					start++
					break
				}
			}
			continue
		}

		if end-start+1 == len(match) {
			startW := words[start]
			counter[startW] = counter[startW] - 1
			result = append(result, start)
			start++
		}
	}

	return result
}

func findSubstring(s string, words []string) []int {
	step := len(words[0])
	result := []int{}

	for offset := 0; offset < step; offset++ {
		newInput := make([]string, 0, (len(s)+2)/3)
		for i := offset; i < len(s); i += step {
			if i+step > len(s) {
				break
			}
			newInput = append(newInput, s[i:i+step])
		}

		tmpResult := findConcatSubstrings(newInput, words)
		for _, index := range tmpResult {
			result = append(result, index*step+offset)
		}
	}
	return result
}
