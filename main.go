package main

func longestCommonPrefix(strs []string) string {
	end := 0
Outer:
	for {
		index := end
		if index >= len(strs[0]) {
			break
		}
		firstCh := strs[0][index]
		for _, s := range strs[1:] {
			if index >= len(s) {
				break Outer
			}
			if s[index] != firstCh {
				break Outer
			}
		}
		end++
	}
	return strs[0][:end]
}
