package Solution

func Solution(n int) int {
	ans, count := 0, 0
	first := true
	for ; n > 0; n >>= 1 {
		if n&1 == 0 {
			if !first {
				count++
			}
			continue
		}
		if first {
			first = false
		} else {
			ans = max(ans, count)
		}
		count = 1
	}
	return ans
}
