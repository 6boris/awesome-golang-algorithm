package Solution

import (
	"sort"
	"strings"
)

func Solution(s string) string {
	array := strings.Split(s, " ")
	sort.Slice(array, func(i, j int) bool {
		a, b := array[i], array[j]
		la, lb := a[len(a)-1], b[len(b)-1]
		return la < lb
	})
	buf := strings.Builder{}
	for i, str := range array {
		buf.WriteString(str[:len(str)-1])
		if i != len(array)-1 {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}
