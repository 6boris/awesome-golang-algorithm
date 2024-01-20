package Solution

import (
	"fmt"
	"strings"
)

type complexNumber struct {
	r, i int
}

func parseComplexNumber(s string) complexNumber {
	cm := complexNumber{}
	idx := strings.Index(s, "+")
	convert := false
	for j := 0; j < idx; j++ {
		if s[j] == '-' {
			convert = true
			continue
		}
		cm.r = cm.r*10 + int(s[j]-'0')
	}
	if convert {
		cm.r = -cm.r
	}
	convert = false
	for j := idx + 1; j < len(s)-1; j++ {
		if s[j] == '-' {
			convert = true
			continue
		}
		cm.i = cm.i*10 + int(s[j]-'0')
	}
	if convert {
		cm.i = -cm.i
	}
	return cm
}
func (c complexNumber) String() string {
	return fmt.Sprintf("%d+%di", c.r, c.i)
}

func mul(a, b complexNumber) complexNumber {
	r := a.r*b.r - a.i*b.i
	i := a.r*b.i + a.i*b.r
	return complexNumber{r: r, i: i}
}
func Solution(num1 string, num2 string) string {
	a := parseComplexNumber(num1)
	b := parseComplexNumber(num2)
	return mul(a, b).String()
}
