package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	if s == "" {
		return []string{}
	}

	return countAndTrim(strings.Fields(s), 10)
}

func countAndTrim(fields []string, length int) []string {
	counter := make(map[string]int)
	for _, field := range fields {
		clearedField := strings.ToLower(strings.Trim(field, "-!.,'"))
		if clearedField == "" {
			continue
		}
		counter[clearedField]++
	}

	words := make([]string, 0, length)
	for word := range counter {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if counter[words[i]] == counter[words[j]] {
			return words[i] < words[j]
		}
		return counter[words[i]] > counter[words[j]]
	})

	if len(words) < length {
		return words
	}

	return words[:length]
}
