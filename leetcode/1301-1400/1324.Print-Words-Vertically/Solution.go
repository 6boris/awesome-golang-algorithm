package Solution

import "strings"

func Solution(s string) []string {
	words := strings.Split(s, " ")
	ml := 0
	for _, w := range words {
		ml = max(ml, len(w))
	}
	ans := make([]string, 0)
	buf := strings.Builder{}
	for index := 0; index < ml; index++ {
		tailSpaces := 0
		for _, w := range words {
			if index < len(w) {
				buf.WriteByte(w[index])
				tailSpaces = 0
				continue
			}
			buf.WriteByte(' ')
			tailSpaces++
		}
		cur := buf.String()
		l := len(cur) - tailSpaces
		cur = cur[:l]
		ans = append(ans, cur)
		buf.Reset()
	}
	return ans
}
