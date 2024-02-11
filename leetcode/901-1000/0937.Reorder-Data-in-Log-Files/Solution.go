package Solution

import (
	"sort"
)

func spaceIndex(str string) int {
	i := 0
	for ; i < len(str) && str[i] != ' '; i++ {
	}
	return i
}
func Solution(logs []string) []string {
	l := len(logs)
	s, e := l-1, l
	for ; s >= 0; s-- {
		if r := logs[s][len(logs[s])-1]; r >= '0' && r <= '9' {
			e--
			logs[s], logs[e] = logs[e], logs[s]
		}
	}
	parts := logs[:e]
	sort.Slice(parts, func(i, j int) bool {
		ai := spaceIndex(logs[i])
		bi := spaceIndex(logs[j])
		if logs[i][ai+1:] == logs[j][bi+1:] {
			return logs[i][:ai] < logs[j][:bi]
		}
		return logs[i][ai+1:] < logs[j][bi+1:]
	})
	return logs
}
