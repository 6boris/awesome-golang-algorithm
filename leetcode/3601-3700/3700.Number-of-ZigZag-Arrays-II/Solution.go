package Solution

const MOD = 1000000007

type Matrix [][]int64

func mul(a, b Matrix) Matrix {
	n := len(a)
	m := len(b[0])
	res := make(Matrix, n)
	for i := range res {
		res[i] = make([]int64, m)
	}

	for i := 0; i < n; i++ {
		for k := 0; k < len(a[0]); k++ {
			r := a[i][k]
			if r == 0 {
				continue
			}
			for j := 0; j < m; j++ {
				res[i][j] = (res[i][j] + r*b[k][j]) % MOD
			}
		}
	}
	return res
}

func powMul(base Matrix, exp int64, res Matrix) Matrix {
	for exp > 0 {
		if exp&1 == 1 {
			res = mul(res, base)
		}
		base = mul(base, base)
		exp >>= 1
	}
	return res
}

func Solution(n int, l int, r int) int {
	m := r - l + 1
	if n == 1 {
		return m
	}

	size := 2 * m
	u := make(Matrix, size)
	for i := range u {
		u[i] = make([]int64, size)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			u[i][j+m] = 1
		}
		for j := i + 1; j < m; j++ {
			u[i+m][j] = 1
		}
	}

	dp := make(Matrix, 1)
	dp[0] = make([]int64, size)
	for i := 0; i < size; i++ {
		dp[0][i] = 1
	}

	dp = powMul(u, int64(n-1), dp)

	var ans int64 = 0
	for i := 0; i < size; i++ {
		ans = (ans + dp[0][i]) % MOD
	}

	return int(ans)
}
