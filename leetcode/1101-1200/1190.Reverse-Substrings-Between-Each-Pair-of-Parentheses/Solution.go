package Solution

import (
	"bytes"
)

func Solution(s string) string {
	stack := make([][]byte, 0)
	buf := bytes.NewBuffer([]byte{})
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == '(' {
			dst := make([]byte, len(buf.Bytes()))
			copy(dst, buf.Bytes())

			stack = append(stack, dst)
			buf.Reset()
			continue
		}
		if s[idx] == ')' {
			tmp := buf.Bytes()
			tmpl := len(tmp)
			bs := make([]byte, tmpl)
			buf.Reset()
			for i := tmpl - 1; i >= 0; i-- {
				bs[tmpl-i-1] = tmp[i]
			}

			l := len(stack) - 1
			top := stack[l]
			stack = stack[:l]

			top = append(top, bs...)
			buf.Write(top)
			continue
		}
		buf.WriteByte(s[idx])
	}

	return buf.String()
}
