package Solution

import "math"

func Solution(s string) int {
	l := len(s)
	ac, bc := make([]int, l), make([]int, l)
	b := 0
	for i := 0; i < l; i++ {
		if s[i] == 'b' {
			b++
		}
		bc[i] = b
	}
	a := 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == 'a' {
			a++
		}
		ac[i] = a
	}
	ans := math.MaxInt
	for i := 0; i < l; i++ {
		ans = min(ans, ac[i]+bc[i])
	}
	return ans - 1
}
