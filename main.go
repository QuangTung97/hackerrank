package main

func isPalindrome(a []byte) bool {
	n := len(a)
	mid := len(a) / 2
	for i := 0; i < mid; i++ {
		if a[i] != a[n-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	byteArr := []byte(s)

	maxLen := 0
	var result []byte

	for i := range byteArr {
		for j := i + 1; j <= len(byteArr); j++ {
			subArr := byteArr[i:j]
			if isPalindrome(subArr) {
				if maxLen < len(subArr) {
					maxLen = len(subArr)
					result = subArr
				}
			}
		}
	}

	return string(result)
}
