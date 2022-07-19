package Solution

const mod1175 = 1000000007

func Solution(n int) int {
	primeArray := make([]uint8, n+1)
	primeCount := 0
	for idx := 2; idx <= n; idx++ {
		if primeArray[idx] == 0 {
			primeCount++
			for next := idx + idx; next <= n; next += idx {
				primeArray[next] = 1
			}
		}
	}

	max_ := primeCount
	notPrimeCount := n - primeCount

	if notPrimeCount > max_ {
		max_ = notPrimeCount
	}
	facPrime, facNotPrime := 1, 1
	a := 1
	for idx := 2; idx <= max_; idx++ {
		a = (a * idx) % mod1175
		if idx == primeCount {
			facPrime = a
		}
		if idx == notPrimeCount {
			facNotPrime = a
		}
	}
	return (facPrime * facNotPrime) % mod1175
}
