package Solution

import "strings"

func Solution(s string, knowledge [][]string) string {
	values := make(map[string]string)
	for _, v := range knowledge {
		values[v[0]] = v[1]
	}

	sb := strings.Builder{}
	var key, replace string
	left := -1
	for i, b := range s {
		if b == '(' {
			left = i
			continue
		}
		if b == ')' {
			key = s[left+1 : i]
			replace = "?"
			if v, ok := values[key]; ok {
				replace = v
			}
			left = -1
			sb.WriteString(replace)
			continue
		}
		if left == -1 {
			sb.WriteByte(byte(b))
		}
	}
	return sb.String()
}
