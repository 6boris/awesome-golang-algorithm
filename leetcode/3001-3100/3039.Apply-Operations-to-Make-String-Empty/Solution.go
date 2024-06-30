package Solution

import (
	"sort"
	"strings"
)

func Solution(s string) string {
	count := [26][2]int{}
	mc := 0
	for i, b := range s {
		count[b-'a'][0]++
		count[b-'a'][1] = i
		mc = max(mc, count[b-'a'][0])
	}
	indies := []int{}
	for i := 0; i < 26; i++ {
		if count[i][0] == mc {
			indies = append(indies, count[i][1])
		}
	}
	sort.Ints(indies)
	sb := strings.Builder{}
	for _, i := range indies {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
