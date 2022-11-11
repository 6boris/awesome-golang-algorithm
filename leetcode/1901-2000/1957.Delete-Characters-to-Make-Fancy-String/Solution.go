package Solution

import "strings"

func Solution(s string) string {
	preByte, repeatCount := byte(' '), 0

	sb := strings.Builder{}
	for idx := 0; idx < len(s); idx++ {
		if preByte == ' ' || preByte != s[idx] {
			sb.WriteByte(s[idx])
			preByte = s[idx]
			repeatCount = 1
			continue
		}
		if repeatCount < 2 {
			repeatCount++
			sb.WriteByte(s[idx])
			continue
		}
	}
	return sb.String()
}
