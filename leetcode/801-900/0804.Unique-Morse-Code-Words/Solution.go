package Solution

import "strings"

var byte2str = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func Solution(words []string) int {
	unique := make(map[string]struct{})

	sb := strings.Builder{}
	for _, word := range words {
		for _, b := range word {
			sb.WriteString(byte2str[b-'a'])
		}
		unique[sb.String()] = struct{}{}
		sb.Reset()
	}
	return len(unique)
}
