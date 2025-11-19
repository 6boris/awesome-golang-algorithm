package Solution

import "bytes"

func convert(s string) string {
	base := 0
	for _, b := range s {
		base = base*10 + int(b-'0')
	}
	buf := bytes.NewBuffer([]byte{})
	for base > 0 {
		mod := base & 1
		if mod == 1 {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}
		base >>= 1
	}
	bs := buf.Bytes()
	for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}
	return string(bs)
}

func Solution(date string) string {
	// 0:4, 5:7 8:end
	return convert(date[:4]) + "-" + convert(date[5:7]) + "-" + convert(date[8:])
}
