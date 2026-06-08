package Solution

import (
	"strconv"
)

func addStrings_2(num1, num2 string) string {
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
