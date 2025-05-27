package Solution

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Solution(n int) []string {
	in := make(map[string]struct{})
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			m := gcd(j, i)
			in[fmt.Sprintf("%d/%d", j/m, i/m)] = struct{}{}
		}
	}
	ans := make([]string, 0)
	for i := range in {
		ans = append(ans, i)
	}
	return ans
}
