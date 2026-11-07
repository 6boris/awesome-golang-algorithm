package Solution

func Solution(n int) int {
	var ret int
	s, e := n, 0
	for ; n > 0; n /= 10 {
		e = e*10 + n%10
	}
	isPrime := func(x int) bool {
		if x <= 1 {
			return false
		}
		if x <= 3 {
			return true
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	for i := min(s, e); i <= max(s, e); i++ {
		if isPrime(i) {
			ret += i
		}
	}

	return ret
}
