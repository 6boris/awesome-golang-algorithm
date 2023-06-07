package Solution

import (
	"fmt"
)

func Solution(a int, b int, c int) int {
	aStr := fmt.Sprintf("%b", a)
	bStr := fmt.Sprintf("%b", b)
	cStr := fmt.Sprintf("%b", c)

	la := len(aStr)
	lb := len(bStr)
	need := 0
	ia, ib := la-1, lb-1
	for i := len(cStr) - 1; i >= 0 || ia >= 0 || ib >= 0; i, ia, ib = i-1, ia-1, ib-1 {
		aa, bb, cc := byte('0'), byte('0'), byte('0')
		if ia >= 0 {
			aa = aStr[ia]
		}
		if ib >= 0 {
			bb = bStr[ib]
		}
		if i >= 0 {
			cc = cStr[i]
		}
		if cc == '0' {
			if aa != '0' {
				need++
			}
			if bb != '0' {
				need++
			}
		}
		if cc == '1' {
			if aa == '0' && bb == '0' {
				need += 1
			}
		}
	}
	return need
}
