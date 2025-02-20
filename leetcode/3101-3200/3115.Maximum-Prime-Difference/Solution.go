package Solution

func Solution(nums []int) int {
	notPrimes := [101]bool{}
	notPrimes[1] = true
	for i := 2; i < 101; i++ {
		if !notPrimes[i] {
			for j := i + i; j < 101; j += i {
				notPrimes[j] = true
			}
		}
	}
	m1, m2 := -1, len(nums)+1
	for i, n := range nums {
		if !notPrimes[n] {
			m1 = max(m1, i)
			m2 = min(m2, i)
		}
	}
	return m1 - m2
}
