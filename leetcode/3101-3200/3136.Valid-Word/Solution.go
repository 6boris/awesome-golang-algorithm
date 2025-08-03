package Solution

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}

func Solution(word string) bool {
	if len(word) < 3 {
		return false
	}
	a, c := false, false
	for _, b := range []byte(word) {
		if b >= '0' && b <= '9' {
			continue
		}
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			if isVowel(b) {
				a = true
			} else {
				c = true
			}
			continue
		}
		return false
	}
	return a && c
}
