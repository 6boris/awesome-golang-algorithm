package Solution

func shouldCal(a, b byte) bool {
	if b == '+' || b == '-' {
		return true
	}
	return a == '*' || a == '/'
}
func cal(x, y int, op byte) int {
	if op == '+' {
		return x + y
	}
	if op == '-' {
		return x - y
	}
	if op == '*' {
		return x * y
	}
	return x / y
}

func Solution(s string) int {
	n := -1
	s = s + "+0"
	a := []byte{}
	b := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if n >= 0 {
				b = append(b, n)
			}
			n = -1
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			if n == -1 {
				n = 0
			}
			n = n*10 + int(s[i]-'0')
			continue
		}
		if n >= 0 {
			b = append(b, n)
		}
		n = -1
		if len(a) == 0 {
			a = append(a, s[i])
			continue
		}
		l := len(a)
		for l > 0 {
			top := a[l-1]
			if !shouldCal(top, s[i]) {
				break
			}
			ll := len(b)
			x, y := b[ll-2], b[ll-1]
			res := cal(x, y, top)
			b[ll-2] = res
			b = b[:ll-1]
			a = a[:l-1]
			l = len(a)
		}
		a = append(a, s[i])
	}
	if n >= 0 {
		b = append(b, n)
	}

	if len(a) == 0 {
		return b[0]
	}
	return cal(b[0], b[1], a[0])
}
