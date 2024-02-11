package Solution

func Solution(n int, primes []int) int {
	//2,3,5 Ugly Number Extensions
	indies := make([]int, len(primes))
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		x := ans[indies[0]] * primes[0]
		for j := 1; j < len(indies); j++ {
			x = min(x, ans[indies[j]]*primes[j])
		}

		ans[i] = x
		for k := 0; k < len(indies); k++ {
			if ans[indies[k]]*primes[k] <= x {
				indies[k]++
			}
		}
	}
	return ans[n-1]
}
