package Solution

import "strings"

func Solution(s string, part string) string {
	ls := len(s)
	lp := len(part)

	for ls >= lp {
		index := strings.Index(s, part)
		if index == -1 {
			break
		}
		// a, b, c, d
		// a, b, c,
		s = s[:index] + s[index+lp:]
		ls = len(s)
	}

	return s
}
