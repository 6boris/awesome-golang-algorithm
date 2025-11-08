package Solution

import "bytes"

func Solution(s string, k int) string {
	buf := bytes.NewBuffer([]byte{})
	sum := 0
	for i := 0; i < len(s); i++ {
		if i != 0 && i%k == 0 {
			buf.WriteByte(byte(sum%26 + 'a'))
			sum = int(s[i]-'a')
			continue
		}
		sum += int(s[i] - 'a')
	}
	buf.WriteByte(byte(sum%26 + 'a'))
	return buf.String()
}
