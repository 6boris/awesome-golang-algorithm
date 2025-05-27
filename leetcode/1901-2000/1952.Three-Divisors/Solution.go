package Solution

func Solution(n int) bool {
	c := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			r := n / i
			if r != i {
				return false
			}
			c++
			if c > 1 {
				return false
			}
		}
	}
	return c == 1
}
