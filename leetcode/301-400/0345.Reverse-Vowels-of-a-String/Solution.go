package Solution

func Solution(ss string) string {
	bs := []byte(ss)
	s, e := 0, len(bs)-1
	for s < e {
		shouldExchange := true
		if !isVowels(bs[s]) {
			s++
			shouldExchange = false
		}
		if !isVowels(bs[e]) {
			e--
			shouldExchange = false
		}
		if shouldExchange {
			bs[s], bs[e] = bs[e], bs[s]
			s, e = s+1, e-1
		}
	}
	return string(bs)
}

func isVowels(b byte) bool {
	return b == 'A' || b == 'a' || b == 'e' || b == 'E' || b == 'i' || b == 'I' || b == 'o' || b == 'O' || b == 'u' || b == 'U'
}
