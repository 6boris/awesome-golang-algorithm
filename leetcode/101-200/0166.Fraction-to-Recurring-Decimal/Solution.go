package Solution

import (
	"bytes"
	"fmt"
)

func Solution(numerator int, denominator int) string {
	const maxn = 65536
	var flag bool
	if numerator*denominator < 0 {
		flag = false
	} else {
		flag = true
	}
	if numerator%denominator == 0 {
		return fmt.Sprintf("%d", numerator/denominator)
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	quotient := make([]int, maxn)
	cf := make([]int, maxn)
	cf[0] = numerator % denominator
	quotient[0] = numerator / denominator
	buf := bytes.NewBufferString("")
	if flag {
		buf.WriteString(fmt.Sprintf("%d.", quotient[0]))
	} else {
		buf.WriteString(fmt.Sprintf("-%d.", quotient[0]))
	}
	for i := 0; i < denominator-1; i++ {
		tmp := (cf[i] * 10) % denominator
		quotient[i+1] = (cf[i] * 10) / denominator
		if tmp == 0 {
			for index := 1; index <= i+1; index++ {
				buf.WriteString(fmt.Sprintf("%d", quotient[index]))
			}
			return buf.String()
		}
		cf[i+1] = tmp
		for index := 0; index <= i; index++ {
			if cf[i+1] == cf[index] {
				if index != 0 {
					for inner := 1; inner <= index; inner++ {
						buf.WriteString(fmt.Sprintf("%d", quotient[inner]))
					}
					buf.WriteString("(")
					for inner := index + 1; inner <= i+1; inner++ {
						buf.WriteString(fmt.Sprintf("%d", quotient[inner]))
					}
					buf.WriteString(")")
				} else {
					buf.WriteString("(")
					for inner := 1; inner <= i+1; inner++ {
						buf.WriteString(fmt.Sprintf("%d", quotient[inner]))
					}
					buf.WriteString(")")
				}
				return buf.String()
			}
		}
	}
	return ""
}
