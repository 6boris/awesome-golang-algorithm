package Solution

func Solution(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	cp1 := make([]int, len(digits))
	copy(cp1, digits)
	i := len(digits) - 1
	remap := -1
	for ; i >= 0; i-- {
		if cp1[i] != 9 {
			remap = cp1[i]
			break
		}
	}
	for ; i >= 0; i-- {
		if cp1[i] == remap {
			cp1[i] = 9
		}
	}

	i = len(digits) - 1
	remap = digits[i]
	for ; i >= 0; i-- {
		if digits[i] == remap {
			digits[i] = 0
		}
	}
	a, b := 0, 0
	for i := len(digits) - 1; i >= 0; i-- {
		a = a*10 + cp1[i]
		b = b*10 + digits[i]
	}
	return a - b
}
