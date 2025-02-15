package Solution

func ok(n, target int) bool {
	if n == target {
		return true
	}
	base := 10
	for ; base <= n; base *= 10 {
		mod := n % base
		if ok(n/base, target-mod) {
			return true
		}
	}
	return false
}
func Solution(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		x := i * i
		res := ok(x, i)
		if res {
			ans += x
		}
	}
	return ans
}
