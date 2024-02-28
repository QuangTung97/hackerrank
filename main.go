package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	last := countAndSay(n - 1)
	var res strings.Builder

	prevCh := last[0]
	count := 1
	for i := 1; i < len(last); i++ {
		ch := last[i]
		if ch == prevCh {
			count++
			continue
		}

		res.WriteString(strconv.FormatInt(int64(count), 10))
		res.WriteByte(prevCh)

		count = 1
		prevCh = ch
	}

	res.WriteString(strconv.FormatInt(int64(count), 10))
	res.WriteByte(prevCh)

	return res.String()
}
