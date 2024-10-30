package Solution

func Solution(n int) int {
	ans := 0
	first := true
	for ; n > 0; n -= 4 {
		cur := n
		if n-1 >= 1 {
			cur *= (n - 1)
		}
		if n-2 >= 1 {
			cur /= (n - 2)
		}
		if n-3 >= 1 {
			if first {
				cur += n - 3
			} else {
				cur -= n - 3
			}
		}
		if first {
			ans = cur
			first = false
			continue
		}
		ans -= cur
	}
	return ans
}
