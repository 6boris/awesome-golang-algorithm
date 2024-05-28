package Solution

func Solution(s string, t string, maxCost int) int {
	start, end := 0, 0
	ans, cur, diff := 0, 0, 0
	for ; end < len(s); end++ {
		if s[end] > t[end] {
			diff = int(s[end] - t[end])
		} else {
			diff = int(t[end] - s[end])
		}
		cur += diff
		for cur > maxCost && start <= end {
			if s[start] > t[start] {
				diff = int(s[start] - t[start])
			} else {
				diff = int(t[start] - s[start])
			}
			cur -= diff
			start++
		}
		ans = max(ans, end-start+1)
	}
	return ans
}
