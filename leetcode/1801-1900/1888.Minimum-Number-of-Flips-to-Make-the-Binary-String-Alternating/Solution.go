package Solution

func Solution(s string) int {
	// Characteristic function
	I := func(ch byte, x int) int {
		if int(ch-'0') == x {
			return 1
		}
		return 0
	}

	n := len(s)
	// Using slice storage
	pre := make([][2]int, n)
	// Note the boundary case when i=0
	for i := 0; i < n; i++ {
		if i == 0 {
			pre[i][0] = I(s[i], 1)
			pre[i][1] = I(s[i], 0)
		} else {
			pre[i][0] = pre[i-1][1] + I(s[i], 1)
			pre[i][1] = pre[i-1][0] + I(s[i], 0)
		}
	}

	ans := min(pre[n-1][0], pre[n-1][1])
	// If the length is odd, consider the move operation
	if n%2 == 1 {
		// If n is an odd number, it is also necessary to calculate suf
		suf := make([][2]int, n)
		// Note the boundary case when i = n - 1
		for i := n - 1; i >= 0; i-- {
			if i == n-1 {
				suf[i][0] = I(s[i], 1)
				suf[i][1] = I(s[i], 0)
			} else {
				suf[i][0] = suf[i+1][1] + I(s[i], 1)
				suf[i][1] = suf[i+1][0] + I(s[i], 0)
			}
		}

		for i := 0; i+1 < n; i++ {
			ans = min(ans, pre[i][0]+suf[i+1][0])
			ans = min(ans, pre[i][1]+suf[i+1][1])
		}
	}

	return ans
}
