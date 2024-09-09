package Solution

import "strings"

func Solution(s string) string {
	left := 1
	leftIndex := 0
	ans := strings.Builder{}
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			left++
			continue
		}
		left--
		if left == 0 {
			ans.WriteString(s[leftIndex+1 : i])
			leftIndex = i + 1
			left = 1
			i++
		}
	}
	return ans.String()
}
