package Solution

import "strings"

func Solution(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	b2w := make(map[byte]string)
	w2b := make(map[string]byte)
	b, w := pattern[0], words[0]
	b2w[b] = w
	w2b[w] = b

	for i := 1; i < len(pattern); i++ {
		if pattern[i] == b {
			if words[i] != w {
				return false
			}
			continue
		}
		b, w = pattern[i], words[i]
		v, ok := b2w[pattern[i]]
		v1, ok1 := w2b[words[i]]
		if (ok && words[i] != v) || (ok1 && pattern[i] != v1) {
			return false
		}
		b2w[b] = w
		w2b[w] = b
	}
	return true
}
