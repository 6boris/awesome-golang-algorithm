package Solution

func Power(b float64, n int) float64 {
	if n < 0 {
		b = 1 / b
		n = -n
	}
	ans := 1.0
	for i := 0; i < n; i++ {
		ans *= b
	}

	return ans
}
