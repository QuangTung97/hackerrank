package main

import (
	"strings"
)

type token struct {
	name string
}

func findNextSlash(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			return s[:i]
		}
	}
	return s
}

func convertToTokens(path string) []token {
	var result []token
	for i := 0; i < len(path); {
		if path[i] == '/' {
			i++
			continue
		}
		name := findNextSlash(path[i:])
		i += len(name)
		result = append(result, token{
			name: name,
		})
	}
	return result
}

func simplifyPath(path string) string {
	tokens := convertToTokens(path)

	var newTokens []token
	for _, tk := range tokens {
		if tk.name == "." {
			continue
		}
		if tk.name == ".." {
			if len(newTokens) > 0 {
				newTokens = newTokens[:len(newTokens)-1]
			}
			continue
		}
		newTokens = append(newTokens, tk)
	}

	if len(newTokens) == 0 {
		return "/"
	}

	var buf strings.Builder
	for _, tk := range newTokens {
		buf.WriteString("/")
		buf.WriteString(tk.name)
	}
	return buf.String()
}
