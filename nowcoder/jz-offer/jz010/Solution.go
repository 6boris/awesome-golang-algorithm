package Solution

//	递推公式
func rectCover(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 0, 2, 0
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
