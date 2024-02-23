package main

func appendStr(a string, b []string) []string {
	result := make([]string, 0, len(b))
	for _, e := range b {
		result = append(result, a+e)
	}
	return result
}

func genRecur(stackSize int, i int, n int) []string {
	if i >= 2*n {
		return []string{""}
	}

	remain := 2*n - i - 1

	var result []string
	if stackSize < n && stackSize < remain {
		result = append(result, appendStr("(", genRecur(stackSize+1, i+1, n))...)
	}
	if stackSize > 0 {
		result = append(result, appendStr(")", genRecur(stackSize-1, i+1, n))...)
	}

	return result
}

func generateParenthesis(n int) []string {
	return genRecur(0, 0, n)
}
