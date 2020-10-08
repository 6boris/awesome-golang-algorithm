package Solution

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	pos := make([]int64, len(num1)+len(num2))
	ans := ""

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1, p2 := i+j, i+j+1
			sum := int64(mul) + pos[p2]

			pos[p1] += sum / 10
			pos[p2] = sum % 10

		}
	}

	for _, v := range pos {
		ans += strconv.Itoa(int(v))
	}
	ans = strings.TrimLeft(ans, "0")
	if len(ans) == 0 {
		return "0"
	}

	return ans
}
