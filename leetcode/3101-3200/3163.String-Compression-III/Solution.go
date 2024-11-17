package Solution

import "strings"

func Solution(word string) string {
	buf := strings.Builder{}
	count := uint8(1)
	bs := []byte(word)
	pre := bs[0]
	for i := 1; i < len(bs); i++ {
		b := bs[i]
		if b == pre {
			count++
			if count == 9 {
				buf.WriteByte(count + '0')
				buf.WriteByte(b)
				count = 0
			}
			continue
		}
		if count != 0 {
			buf.WriteByte(count + '0')
			buf.WriteByte(pre)
		}
		pre = b
		count = 1
	}
	if count != 0 {
		buf.WriteByte(count + '0')
		buf.WriteByte(pre)
	}
	return buf.String()
}
