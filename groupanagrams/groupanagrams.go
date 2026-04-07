package groudanagrams

import (
	"sort"
)

func GroupAnagrams(words []string) map[string][]string {
	m := make(map[string][]string)
	for _, word := range words {
		key := sortWord(word)
		m[key] = append(m[key], word)
	}

	return m
}

func sortWord(word string) string {
	b := []byte(word)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
