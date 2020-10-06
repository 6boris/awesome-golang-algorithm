package Solution

import "math"

//	递归题解
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//	自顶向下推导
func Fibonacci2(n int) int {
	a, b := 0, 1

	for i := 1; i <= n; i++ {
		a = a + b
		b = a - b
	}
	return a
}

//	推导公式1
func Fibonacci3(n int) int {
	a := math.Sqrt(5) / 5
	b := math.Pow((1+math.Sqrt(5))/2, float64(n))
	c := math.Pow((1-math.Sqrt(5))/2, float64(n))
	return int(a * (b - c))
}
