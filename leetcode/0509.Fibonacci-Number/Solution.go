package Solution

func Solution(N int) int {
	var ans int
	fibSeq := getFibSeq()
	for i := 0; i <= N; i++ {
		ans = fibSeq()
	}
	return ans
}

func getFibSeq() func() int {
	a, b := -1, 1
	return func() int {
		c := a + b
		a, b = b, c
		return c
	}
}
