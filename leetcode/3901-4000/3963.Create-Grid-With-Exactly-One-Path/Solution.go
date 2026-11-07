package Solution

import "strings"

func Solution(m int, n int) []string {
	ret := make([]string, m)
	var buf strings.Builder
	buf.WriteByte('.')
	for i := 1; i < n; i++ {
		buf.WriteByte('#')
	}
	path := buf.String()
	for i := range m - 1 {
		ret[i] = path
	}
	buf.Reset()
	for range n {
		buf.WriteByte('.')
	}
	ret[m-1] = buf.String()

	return ret
}
