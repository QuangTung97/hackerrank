package main

import (
	"strings"
)

type symbol struct {
	str string
	val int
}

var symbols = []symbol{
	{str: "I", val: 1},
	{str: "IV", val: 4},
	{str: "V", val: 5},
	{str: "IX", val: 9},
	{str: "X", val: 10},
	{str: "XL", val: 40},
	{str: "L", val: 50},
	{str: "XC", val: 90},
	{str: "C", val: 100},
	{str: "CD", val: 400},
	{str: "D", val: 500},
	{str: "CM", val: 900},
	{str: "M", val: 1000},
}

func intToRoman(num int) string {
	var b strings.Builder
Outer:
	for num > 0 {
		for i := len(symbols) - 1; i >= 0; i-- {
			s := symbols[i]
			if num >= s.val {
				num -= s.val
				b.WriteString(s.str)
				continue Outer
			}
		}
	}
	return b.String()
}
