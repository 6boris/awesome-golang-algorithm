package Solution

func binom(a, b int) int64 {
	n, k := int64(a), int64(b)
	if n < 0 || k < 0 || n < k {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	var result int64 = 1
	for i := int64(1); i <= k; i++ {
		result = result * (n - i + 1) / i
	}
	return result
}

func Solution(n int, limit int) int64 {
	total := binom(n+2, 2)

	n1 := n - (limit + 1)
	n2 := n - 2*(limit+1)
	n3 := n - 3*(limit+1)

	total -= 3 * binom(n1+2, 2)
	total += 3 * binom(n2+2, 2)
	total -= binom(n3+2, 2)

	return total
}
