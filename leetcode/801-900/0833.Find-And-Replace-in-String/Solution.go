package Solution

import (
	"sort"
	"strings"
)

func Solution(s string, indices []int, sources []string, targets []string) string {
	position := make([]int, len(indices))
	for i := 0; i < len(indices); i++ {
		position[i] = i
	}
	sort.Slice(position, func(i, j int) bool {
		return indices[position[i]] < indices[position[j]]
	})

	length := len(s)
	stack := make([]int, len(indices))
	cur := -1
	for _, pos := range position {
		sIndex := indices[pos]
		if left := length - sIndex; left < len(sources[pos]) {
			continue
		}
		if s[sIndex:sIndex+len(sources[pos])] == sources[pos] {
			if cur == -1 {
				cur++
				stack[cur] = pos
				continue
			}
			lastestPos := stack[cur]
			if sIndex >= indices[lastestPos]+len(sources[lastestPos]) {
				cur++
				stack[cur] = pos
				continue
			}
			cur--
		}
	}
	if cur == -1 {
		return s
	}
	start := 0
	buf := strings.Builder{}
	for i := 0; i < cur+1; i++ {
		pos := stack[i]
		sIndex := indices[pos]
		if sIndex > start {
			buf.WriteString(s[start:sIndex])
		}
		buf.WriteString(targets[pos])
		start = sIndex + len(sources[pos])
	}
	if start < len(s) {
		buf.WriteString(s[start:])
	}
	return buf.String()
}
