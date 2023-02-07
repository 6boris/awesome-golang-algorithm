package Solution

func Solution(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	// t0, t1, t2
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		x := a + b + c
		a, b, c = b, c, x
	}
	return c
}
