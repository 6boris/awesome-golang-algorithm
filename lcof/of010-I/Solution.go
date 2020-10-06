package Solution

import "math"

//	递归题解
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return (fib(n-1) + fib(n-2)) % 1000000007
}

//	自顶向下推导
func fib2(n int) int {
	a, b, sum := 0, 1, 0

	for i := 1; i <= n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}
	return a
}

//	推导公式1
func fib3(n int) int {
	a := math.Sqrt(5) / 5
	b := math.Pow((1+math.Sqrt(5))/2, float64(n))
	c := math.Pow((1-math.Sqrt(5))/2, float64(n))
	return int(a*(b-c)) % 1000000007
}
