package Solution

import (
	"bytes"
)

type stackItem struct {
	num int
	str string
}

func Solution(s string) string {
	l := len(s)
	ans := bytes.NewBuffer([]byte{})

	stack := make([]stackItem, 0)
	buf := bytes.NewBuffer([]byte{})
	base := 0
	// 3[z] 2[2[y]pq4[2[jk]e1[f]]]ef"
	//        2
	// zzz    yypg jkjkef jkjkef jkjkef jkjkef yypg jkjkef jkjkef jkjkef jkjkef  ef
	for i := 0; i < l; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			base = 0
			if len(stack) == 0 {
				ans.WriteByte(s[i])
			} else {
				buf.WriteByte(s[i])
			}
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			base = base*10 + int(s[i]-'0')
			r := buf.String()
			if ls := len(stack); ls > 0 {
				stack[ls-1].str += r
			}
			buf.Reset()
			continue
		}
		if s[i] == '[' {
			stack = append(stack, stackItem{num: base})
			base = 0
			continue
		}
		ls := len(stack)
		top := stack[ls-1]
		stack = stack[:ls-1]
		r := buf.String()
		top.str += r
		buf.Reset()
		for i := 0; i < top.num; i++ {
			buf.WriteString(top.str)
		}
		r = buf.String()
		if ls := len(stack); ls > 0 {
			stack[ls-1].str += r
		} else {
			ans.WriteString(r)
		}
		buf.Reset()
	}
	return ans.String()
}
