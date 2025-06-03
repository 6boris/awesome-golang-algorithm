package Solution

func binom(a, b int) int {
	n, k := a, b
	if n < 0 || k < 0 || n < k {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	var result int = 1
	for i := 1; i <= k; i++ {
		result = result * (n - i + 1) / i
	}
	return result
}

func Solution(n int, limit int) int {
	total := binom(n+2, 2)

	n1 := n - (limit + 1)
	n2 := n - 2*(limit+1)
	n3 := n - 3*(limit+1)

	total -= 3 * binom(n1+2, 2)
	total += 3 * binom(n2+2, 2)
	total -= binom(n3+2, 2)

	return total
}
