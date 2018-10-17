package Solution

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / pow(x, -n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	ans := pow(x, n>>1)
	if n&1 == 0 {
		return ans * ans
	}
	return ans * ans * x
}
