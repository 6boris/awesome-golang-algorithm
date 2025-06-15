package Solution

func Solution(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	cp := make([]int, len(digits))
	copy(cp, digits)

	i := len(digits) - 1
	remap := -1
	for ; i >= 0; i-- {
		if cp[i] != 9 {
			remap = cp[i]
			break
		}
	}
	for ; i >= 0; i-- {
		if cp[i] == remap {
			cp[i] = 9
		}
	}

	i = len(digits) - 1
	target := 1
	if digits[i] == 1 {
		target = 0
	}

	remap = -1
	for ; i >= 0; i-- {
		if digits[i] == 0 {
			continue
		}
		if digits[i] != 1 {
			remap = digits[i]
			break
		}
	}
	for ; i >= 0; i-- {
		if digits[i] == remap {
			digits[i] = target
		}
	}
	a, b := 0, 0
	for i := len(digits) - 1; i >= 0; i-- {
		a = a*10 + cp[i]
		b = b*10 + digits[i]
	}
	return a - b
}
