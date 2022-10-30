package Solution

import "strings"

func Solution(word1 []string, word2 []string) bool {
	sb1, sb2 := strings.Builder{}, strings.Builder{}

	for _, s := range word1 {
		sb1.WriteString(s)
	}

	for _, s := range word2 {
		sb2.WriteString(s)
	}

	return sb1.String() == sb2.String()
}
