package Solution

func romanToInt(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	sum := m[string(s[len(s)-1])]
	// 从后向前遍历
	// 每次和前面一位数比较
	for i := len(s) - 2; i >= 0; i-- {
		if m[string(s[i])] < m[string(s[i+1])] {
			sum -= m[string(s[i])]
		} else {
			sum += m[string(s[i])]
		}
	}

	return sum
}

func romanToInt2(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	ans := 0
	// 从前向后遍历
	// 每次默认相加再检查和前面一位数的大小
	// 前面 > 后面  ans = s[i-1] + s[i]
	// 前面 < 后面  ans = - (2 * s[i-1]) + s[i] (因为默认加了一次所以需要减2次)
	for i := range s {
		ans += m[string(s[i])]
		if i > 0 && m[string(s[i])] > m[string(s[i-1])] {
			ans -= 2 * m[string(s[i-1])]
		}
	}
	return ans
}
