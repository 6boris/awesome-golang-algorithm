package Solution

import "strings"

func Solution(s string, p string) bool {
	pos := -1
	size := len(p)
	for i := range size {
		if p[i] == '*' {
			pos = i
			break
		}
	}

	if pos == -1 {
		return strings.Index(s, p) != -1
	}
	if pos == 0 {
		p = p[1:]
	}
	if pos == size-1 {
		p = p[:pos]
	}
	if len(p) != size {
		return strings.Index(s, p) != -1
	}

	pre := p[:pos]
	suf := p[pos+1:]
	pi := strings.Index(s, pre)
	if pi == -1 {
		return false
	}
	si := strings.LastIndex(s, suf)
	if si == -1 {
		return false
	}
	// ab cd
	// 01 23
	return si >= pi+len(pre)
}
