package Solution

func Solution(s string, k int, fill byte) []string {
	l := len(s)
	bs := []byte(s)
	mod := l % k
	if mod != 0 {
		for i := 0; i < k-mod; i++ {
			bs = append(bs, fill)
		}
	}
	var ans []string
	for i := 0; i < len(bs); i += k {
		ans = append(ans, string(bs[i:i+k]))
	}
	return ans
}
