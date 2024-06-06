package Solution

import "strings"

func Solution(s string) string {
	left := -1
	stack := make([]int, len(s))
	removes := make([]int, 0)
	for i := range s {
		if s[i] == '(' {
			left++
			stack[left] = i
			continue
		}
		if s[i] == ')' {
			if left >= 0 {
				left--
			} else {
				removes = append(removes, i)
			}
		}
	}
	check := make(map[int]struct{})
	for i := 0; i <= left; i++ {
		check[stack[i]] = struct{}{}
	}
	for _, idx := range removes {
		check[idx] = struct{}{}
	}
	buf := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if _, ok := check[i]; ok {
			continue
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
