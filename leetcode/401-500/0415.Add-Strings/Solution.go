package Solution

import "bytes"

func Solution(num1 string, num2 string) string {
	num1Idx, num2Idx := len(num1)-1, len(num2)-1

	buf := bytes.NewBufferString("")
	cf := uint8(0)
	for ; num1Idx >= 0 && num2Idx >= 0; num1Idx, num2Idx = num1Idx-1, num2Idx-1 {
		sum := num1[num1Idx] - '0' + num2[num2Idx] - '0' + cf
		cf = sum / 10
		sum %= 10
		buf.WriteByte(sum + '0')
	}

	for ; num1Idx >= 0; num1Idx-- {
		sum := num1[num1Idx] - '0' + cf
		cf = sum / 10
		sum %= 10
		buf.WriteByte(sum + '0')
	}
	for ; num2Idx >= 0; num2Idx-- {
		sum := num2[num2Idx] - '0' + cf
		cf = sum / 10
		sum %= 10
		buf.WriteByte(sum + '0')
	}

	if cf != 0 {
		buf.WriteByte(cf + '0')
	}
	bs := buf.Bytes()
	for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}

	return string(bs)
}
