package main

var charMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func appendStr(a string, b []string) []string {
	var result []string
	for _, e := range b {
		result = append(result, a+e)
	}
	return result
}

func handleRecur(digits string) []string {
	if len(digits) == 0 {
		return []string{""}
	}

	ch := digits[0]

	var result []string
	values := charMap[ch]
	for _, v := range values {
		result = append(result,
			appendStr(v, handleRecur(digits[1:]))...,
		)
	}
	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return handleRecur(digits)
}
