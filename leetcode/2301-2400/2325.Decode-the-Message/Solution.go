package Solution

import "bytes"

func Solution(key string, message string) string {
	bucket := make([]uint8, 26) // space

	char := 'a'
	for _, b := range key {
		if b != ' ' && bucket[b-'a'] == 0 {
			bucket[b-'a'] = uint8(char)
			char++
		}
	}
	buf := bytes.NewBuffer([]byte{})
	for _, msg := range message {
		if msg == ' ' {
			buf.WriteByte(byte(msg))
			continue
		}
		buf.WriteByte(bucket[msg-'a'])
	}
	return buf.String()
}
