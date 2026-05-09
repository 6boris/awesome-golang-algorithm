package Solution

import (
	"sort"
	"strings"
)

func Solution(s string) string {
	count := 0
	start := 0
	var parts []string

	for i, ch := range s {
		if ch == '1' {
			count++
		} else {
			count--
		}

		if count == 0 {
			inner := s[start+1 : i]
			innerMax := Solution(inner)
			parts = append(parts, "1"+innerMax+"0")
			start = i + 1
		}
	}

	sort.Slice(parts, func(i, j int) bool {
		return parts[i] > parts[j]
	})

	var buf strings.Builder
	for i := range parts {
		buf.WriteString(parts[i])
	}
	return buf.String()
}
