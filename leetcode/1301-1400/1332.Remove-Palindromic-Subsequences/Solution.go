package Solution

func Solution(s string) int {
	for s1, e := 0, len(s)-1; s1 < e; s1, e = s1+1, e-1 {
		if s[s1] != s[e] {
			return 2
		}
	}
	return 1
}
