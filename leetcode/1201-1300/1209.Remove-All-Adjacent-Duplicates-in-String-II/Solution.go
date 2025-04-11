package Solution

import "strings"

type item1209 struct {
	b byte
	c int
}

func Solution(s string, k int) string {
	stack := make([]item1209, len(s))
	index := -1
	for i := 0; i < len(s); i++ {
		if index == -1 {
			index++
			stack[index] = item1209{b: s[i], c: 1}
			continue
		}
		if s[i] == stack[index].b {
			stack[index].c++
			if stack[index].c == k {
				index--
			}
			continue
		}
		index++
		stack[index] = item1209{b: s[i], c: 1}
	}
	buf := strings.Builder{}
	for i := 0; i <= index; i++ {
		for ; stack[i].c > 0; stack[i].c-- {
			buf.WriteByte(stack[i].b)
		}
	}

	return buf.String()
}
