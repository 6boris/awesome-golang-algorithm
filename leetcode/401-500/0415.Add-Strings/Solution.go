package Solution

import (
	"bytes"
	"strconv"
)

func addStrings_1(num1 string, num2 string) string {
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

func addStrings_2(num1 string, num2 string) string {
	ans, carry := "", 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		tmp := x + y + carry
		ans = strconv.Itoa(tmp%10) + ans
		carry = tmp / 10
	}
	return ans
}
