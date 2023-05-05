package Solution

func isVowel(x byte) bool {
	return x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u'
}
func Solution(s string, k int) int {
	ans := 0
	start, end := 0-k, 0
	vowels := 0
	for ; end < len(s); end++ {
		if isVowel(s[end]) {
			vowels++
		}
		if start >= 0 && isVowel(s[start]) {
			vowels--
		}

		if start >= -1 && vowels > ans {
			ans = vowels
		}

		start++
	}
	return ans
}
