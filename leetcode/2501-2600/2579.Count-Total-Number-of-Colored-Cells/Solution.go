package Solution

func Solution(n int) int64 {
	ans := int64(1)
	var add int64
	for i := 2; i <= n; i++ {
		add = int64(i-1) * 4
		ans += add
	}

	return ans
}
