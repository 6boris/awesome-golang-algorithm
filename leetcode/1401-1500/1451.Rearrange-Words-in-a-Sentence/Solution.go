package Solution

import (
	"sort"
	"strings"
)

func Solution(text string) string {
	if len(text) == 0 {
		return ""
	}
	bs := strings.Split(text, " ")
	bs[0] = strings.ToLower(bs[0])

	indies := make([]int, len(bs))
	for i := 0; i < len(bs); i++ {
		indies[i] = i
	}
	sort.Slice(indies, func(i, j int) bool {
		a, b := len(bs[indies[i]]), len(bs[indies[j]])
		if a == b {
			return indies[i] < indies[j]
		}
		return a < b
	})
	ans := make([]string, len(bs))
	for i := 0; i < len(bs); i++ {
		ans[i] = bs[indies[i]]
	}
	first := []byte(ans[0])
	first[0] -= 32
	ans[0] = string(first)
	return strings.Join(ans, " ")
}
