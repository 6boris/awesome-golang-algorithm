package Solution

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	wc, l := 0, len(s)-1
	for i := l; i >= 0; i-- {
		if s[i] != ' ' {
			wc++
		} else if wc > 0 {
			return wc
		}
	}
	return wc
}

func lengthOfLastWord1(s string) int {
	s = strings.TrimRight(s, " ")
	return len(s) - strings.LastIndex(s, " ") - 1
}
