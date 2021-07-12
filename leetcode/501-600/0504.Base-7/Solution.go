package Solution

import "bytes"

func Solution(num int) string {
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}

	bf := bytes.NewBufferString("")
	for num >= 7 {
		mod := num % 7
		num /= 7
		bf.WriteByte(byte(mod) + 48)
	}

	bf.WriteByte(byte(num) + 48)
	if negative {
		bf.WriteByte('-')
	}

	bs := bf.Bytes()
	for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}
	return string(bs)
}
