package Solution

const mod1922 = 1000000007

func quickPow(a int, n int64) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}

	ans := quickPow(a, n>>1)
	ans = (ans * ans) % mod1922
	if n&1 == 1 {
		ans = (ans * a) % mod1922
	}
	return ans
}

func Solution(n int64) int {
	add := int64(0)
	if n&1 == 1 {
		add = 1
	}
	mid := n >> 1
	return (quickPow(5, mid+add) * quickPow(4, mid)) % mod1922
}
