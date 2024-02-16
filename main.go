package main

type bitset struct {
	data [2]uint64
}

const shift = 6
const mask = (1 << shift) - 1

func (b *bitset) add(input int) {
	x := uint64(input)
	outer := x >> shift
	offset := x & mask
	oneHot := uint64(1 << offset)
	b.data[outer] |= oneHot
}

func (b *bitset) existed(input int) bool {
	x := uint64(input)
	outer := x >> shift
	offset := x & mask
	oneHot := uint64(1 << offset)
	return (b.data[outer] & oneHot) > 0
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0

	byteArr := []byte(s)
	for i := range byteArr {
		var set bitset
		for j := i; j < len(byteArr); j++ {
			ch := byteArr[j]
			if set.existed(int(ch)) {
				break
			}
			set.add(int(ch))

			newLen := j - i + 1
			if maxLen < newLen {
				maxLen = newLen
			}
		}
	}

	return maxLen
}
