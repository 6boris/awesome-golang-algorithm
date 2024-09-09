package Solution

func Solution(s string, k int) string {
	bs := []byte(s)
	start := 0
	for start < len(s) {
		nextStart := start + 2*k
		for s, e := start, min(start+k-1, len(s)-1); s < e; s, e = s+1, e-1 {
			bs[s], bs[e] = bs[e], bs[s]
		}
		start = nextStart
	}
	return string(bs)
}
