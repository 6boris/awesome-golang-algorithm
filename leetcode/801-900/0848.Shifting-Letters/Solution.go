package Solution

func Solution(s string, shifts []int) string {
	ans := []byte(s)
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		sum += shifts[i]
		now := (int(ans[i]-'a') + sum) % 26
		ans[i] = byte(now) + 'a'
	}
	return string(ans)
}
