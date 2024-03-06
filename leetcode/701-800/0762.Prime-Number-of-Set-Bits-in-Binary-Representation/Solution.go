package Solution

func bitOf1(n int) int {
	c := 0
	for n > 0 {
		c++
		n = n & (n - 1)
	}
	return c
}
func Solution(left int, right int) int {
	// 1, 2, 3,  4,  5,   6, 7,   8,   9,    10
	// 1, 1  11 100  101 110 111 1000  1001  1010
	notPrimes := make([]bool, 33)
	notPrimes[1] = true
	for i := 2; i < 33; i++ {
		if notPrimes[i] {
			continue
		}
		for next := i + i; next < 33; next += i {
			notPrimes[next] = true
		}
	}
	ans := 0
	for i := left; i <= right; i++ {
		if r := bitOf1(i); !notPrimes[r] {
			ans++
		}
	}
	return ans
}
