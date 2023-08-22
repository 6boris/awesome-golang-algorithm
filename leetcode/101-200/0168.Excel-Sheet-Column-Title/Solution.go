package Solution

import "bytes"

func Solution(columnNumber int) string {
	base := 64
	buf := bytes.NewBuffer([]byte{})
	for columnNumber > 0 {
		mod := columnNumber % 26
		columnNumber /= 26
		if mod == 0 {
			mod = 26
			columnNumber--
		}
		buf.WriteByte(uint8(base + mod))
	}

	bs := buf.Bytes()
	for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}
	return string(bs)
}
