package Solution

func jumpFloorII(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 1, 0
	for i := 2; i <= n; i++ {
		b = a << 1
		a = b
	}
	return b
}
