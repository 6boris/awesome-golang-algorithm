package Solution

const mod1513 = 1e9 + 7

func Solution(s string) int {
	one := 0
	ans := 0
	for _, b := range s {
		if b == '0' {
			// 1111
			// 4 3 2 1
			ans = (ans + one*(one+1)/2) % mod1513
			one = 0
			continue
		}
		one++
	}
	ans = (ans + one*(one+1)/2) % mod1513
	return ans
}
