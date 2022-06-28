package Solution

import "strings"

func Solution(s string, k int) string {
	sb := strings.Builder{}
	words := 0
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == ' ' {
			words++
			if words == k {
				break
			}
		}
		sb.WriteByte(s[idx])
	}
	return sb.String()
}
