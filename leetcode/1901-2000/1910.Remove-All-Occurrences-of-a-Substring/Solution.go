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

func Solution1(s string, part string) string {
	bs := []byte(s)
	index := -1
	l := len(part)
	lastByte := part[l-1]
	for i := 0; i < len(s); i++ {
		index++
		bs[index] = s[i]
		if s[i] != lastByte {
			continue
		}

		if start := index - l + 1; start >= 0 {
			if string(bs[start:index+1]) == part {
				index = start - 1
			}
		}

	}
	if start := index - l + 1; start >= 0 {
		if string(bs[start:index+1]) == part {
			index = start - 1
		}
	}

	return string(bs[:index+1])
}
