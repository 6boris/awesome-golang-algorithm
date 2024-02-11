package Solution

import (
	"fmt"
	"strings"
)

func parseExpression(exp string) (int, int) {
	a, b := 0, 0
	idx := 0
	neg := false
	cur := 0
	for ; idx < len(exp); idx++ {
		if exp[idx] >= '0' && exp[idx] <= '9' {
			cur = cur*10 + int(exp[idx]-'0')
			continue
		}

		if exp[idx] == 'x' {
			if cur == 0 && (idx == 0 || exp[idx-1] != '0') {
				cur++
			}
			if neg {
				cur = -cur
			}
			a += cur
			cur = 0
		} else {
			if neg {
				cur = -cur
			}
			b += cur
			cur = 0
			neg = exp[idx] == '-'
		}
	}
	if neg {
		cur = -cur
	}
	b += cur
	return a, b
}

func gcd640(a, b int) int {
	mod := a % b
	for mod != 0 {
		a, b = b, mod
		mod = a % b
	}
	return b
}

func Solution(equation string) string {
	expressions := strings.Split(equation, "=")
	la, lb := parseExpression(expressions[0])
	ra, rb := parseExpression(expressions[1])
	x := la - ra
	y := rb - lb

	if x < 0 {
		x = -x
		y = -y
	}
	if x == 0 {
		if y != 0 {
			return "No solution"
		}
		return "Infinite solutions"
	}
	if y != 0 {
		a, b := x, y
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		c := gcd640(a, b)
		x /= c
		y /= c
	} else {
		x = 1
	}
	if x == 1 {
		return fmt.Sprintf("x=%d", y)
	}
	return fmt.Sprintf("%dx=%d", x, y)
}
