package Solution

const (
	MOD = 1e9 + 7
	MX  = 100000
)

var (
	fact    = make([]int64, MX)
	invFact = make([]int64, MX)
)

func qpow(x int64, n int) int64 {
	res := int64(1)
	for n > 0 {
		if n&1 == 1 {
			res = res * x % MOD
		}
		x = x * x % MOD
		n >>= 1
	}
	return res
}

func initFact() {
	if fact[0] != 0 {
		return
	}
	fact[0] = 1
	for i := 1; i < MX; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}
	invFact[MX-1] = qpow(fact[MX-1], MOD-2)
	for i := MX - 1; i > 0; i-- {
		invFact[i-1] = invFact[i] * int64(i) % MOD
	}
}

func comb(n, m int) int64 {
	return fact[n] * invFact[m] % MOD * invFact[n-m] % MOD
}

func Solution(n, m, k int) int {
	initFact()
	return int(comb(n-1, k) * int64(m) % MOD * qpow(int64(m-1), n-k-1) % MOD)
}
