package Solution

const (
	MOD   int = 1e9 + 7
	MAX_N     = 10010
	MAX_P     = 15 // There are up to 15 prime factors
)

var (
	c     [MAX_N + MAX_P][MAX_P + 1]int
	sieve [MAX_N]int   // Minimum prime factor
	ps    [MAX_N][]int // List of prime factor counts
)

func initialize() {
	if c[0][0] != 0 {
		return
	}

	for i := 2; i < MAX_N; i++ {
		if sieve[i] == 0 {
			for j := i; j < MAX_N; j += i {
				if sieve[j] == 0 {
					sieve[j] = i
				}
			}
		}
	}

	for i := 2; i < MAX_N; i++ {
		x := i
		for x > 1 {
			p := sieve[x]
			cnt := 0
			for x%p == 0 {
				x /= p
				cnt++
			}
			ps[i] = append(ps[i], cnt)
		}
	}

	c[0][0] = 1
	for i := 1; i < MAX_N+MAX_P; i++ {
		c[i][0] = 1
		for j := 1; j <= MAX_P && j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % MOD
		}
	}
}

func Solution(n int, maxValue int) int {
	initialize()
	ans := 0
	for x := 1; x <= maxValue; x++ {
		mul := 1
		for _, p := range ps[x] {
			mul = mul * c[n+p-1][p] % MOD
		}
		ans = (ans + mul) % MOD
	}
	return ans
}
