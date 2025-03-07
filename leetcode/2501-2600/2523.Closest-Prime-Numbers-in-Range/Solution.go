package Solution

import "sort"

func calPrime(max int) []int {
	isPrime := make([]bool, max+1)

	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= max; i++ {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}

	var primes []int
	for i := 2; i <= max; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func Solution(left int, right int) []int {
	res := []int{-1, -1}
	primes := calPrime(right)
	start := sort.Search(len(primes), func(i int) bool {
		return primes[i] >= left
	})
	l := primes[start:]
	m := -1
	for i := 1; i < len(l); i++ {
		if m == -1 || l[i]-l[i-1] < m {
			m = l[i] - l[i-1]
			res[0] = l[i-1]
			res[1] = l[i]
		}
	}
	return res
}
