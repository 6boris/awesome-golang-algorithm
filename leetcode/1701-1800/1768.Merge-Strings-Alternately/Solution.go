package Solution

import "strings"

func Solution(word1 string, word2 string) string {
	sb := strings.Builder{}
	a, b := 0, 0
	for ; a < len(word1) || b < len(word2); a, b = a+1, b+1 {
		if a < len(word1) {
			sb.WriteByte(word1[a])
		}
		if b < len(word2) {
			sb.WriteByte(word2[b])
		}
	}
	return sb.String()
}
