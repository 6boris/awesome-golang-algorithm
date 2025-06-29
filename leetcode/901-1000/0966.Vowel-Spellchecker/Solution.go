package Solution

import (
	"fmt"
	"strings"
)

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}

// babee = b1b2
func toVowel(word string) string {
	buf := strings.Builder{}
	c := 0
	for _, b := range []byte(word) {
		if !isVowel(b) {
			buf.WriteByte(b)
			continue
		}
		buf.WriteString(fmt.Sprintf("%d", c))
		c = 0
	}
	if c > 0 {
		buf.WriteString(fmt.Sprintf("%d", c))
	}

	return buf.String()
}

func Solution(wordlist []string, queries []string) []string {
	m := make(map[string]struct{})
	lower := make(map[string]string)
	matched := make(map[string]string)
	for _, word := range wordlist {
		m[word] = struct{}{}
		l := strings.ToLower(word)
		if _, ok := lower[l]; !ok {
			lower[l] = word
		}
		match := toVowel(l)
		if _, ok := matched[match]; !ok {
			matched[match] = word
		}
	}
	ans := make([]string, len(queries))
	for i, q := range queries {
		if _, ok := m[q]; ok {
			ans[i] = q
			continue
		}
		lq := strings.ToLower(q)

		if v, ok := lower[lq]; ok {
			ans[i] = v
			continue
		}
		if v, ok := matched[toVowel(lq)]; ok {
			ans[i] = v
			continue
		}
	}
	return ans
}
