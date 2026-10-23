package Solution

func Solution(n int, k int) int {
	const mod = 1000000007

	m := n + k - 1
	r := 2 * k

	if r > m || r < 0 {
		return 0
	}

	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0] = 1
		c[i][i] = 1
		for j := 1; j < i; j++ {
			c[i][j] = (c[i-1][j-1] + c[i-1][j]) % mod
		}
	}

	return c[m][r]
}
