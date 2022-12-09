package Solution

func Solution(s string) bool {
	l, r := 0, 0
	// ab
	for s1, e := 0, len(s)-1; s1 < e; s1, e = s1+1, e-1 {
		if isVowels1704(s[s1]) {
			l++
		}
		if isVowels1704(s[e]) {
			r++
		}
	}
	return l == r
}
func isVowels1704(b byte) bool {
	return b == 'a' || b == 'A' || b == 'e' || b == 'E' || b == 'i' || b == 'I' || b == 'o' || b == 'O' || b == 'u' || b == 'U'
}
