package Solution

import "fmt"

func gcd592(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func add592(a, b, c, d int) (int, int) {
	denominator := b * d
	numerator := a*d + b*c
	lessZero := false
	if numerator < 0 {
		lessZero = true
		numerator = -numerator
	}
	g := gcd592(numerator, denominator)
	denominator /= g
	numerator /= g
	if lessZero {
		numerator = -numerator
	}

	return numerator, denominator
}
func sub592(a, b, c, d int) (int, int) {

	denominator := b * d
	numerator := a*d - b*c
	lessZero := false
	if numerator < 0 {
		lessZero = true
		numerator = -numerator
	}
	g := gcd592(numerator, denominator)
	denominator /= g
	numerator /= g
	if lessZero {
		numerator = -numerator
	}

	return numerator, denominator
}

func Solution(expression string) string {
	numerator, denominator := 0, 1
	nums := 0
	op := byte('+')
	if expression[0] != '-' {
		nums++
	}
	n, d := 0, 0
	f := true
	for index := 0; index < len(expression); index++ {
		if expression[index] == '+' || expression[index] == '-' {
			nums++
			if nums == 2 {
				if op == '+' {
					numerator, denominator = add592(numerator, denominator, n, d)
				} else {
					numerator, denominator = sub592(numerator, denominator, n, d)
				}
				nums = 1
			}

			op = expression[index]
			f = true
			n, d = 0, 0
			continue
		}
		if expression[index] == '/' {
			f = false
		}
		if expression[index] >= '0' && expression[index] <= '9' {
			if f {
				n = n*10 + int(expression[index]-'0')
			} else {
				d = d*10 + int(expression[index]-'0')
			}
		}
	}
	if op == '+' {
		numerator, denominator = add592(numerator, denominator, n, d)
	} else {
		numerator, denominator = sub592(numerator, denominator, n, d)
	}
	return fmt.Sprintf("%d/%d", numerator, denominator)
}
