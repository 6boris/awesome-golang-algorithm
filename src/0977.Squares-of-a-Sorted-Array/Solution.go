package Solution

import "math"

func Solution(A []int) []int {
	n := len(A)
	if n == 0 {
		return A
	}
	abs := func(num int) int { return int(math.Abs(float64(num))) }
	ans := make([]int, n)
	i, j, k := 0, n-1, n-1
	for i <= j {
		if abs(A[i]) > abs(A[j]) {
			ans[k] = A[i] * A[i]
			i++
		} else {
			ans[k] = A[j] * A[j]
			j--
		}
		k--
	}
	return ans
}
