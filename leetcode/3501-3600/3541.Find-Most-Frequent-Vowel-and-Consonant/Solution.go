package Solution

func Solution(s string) int {
	count := [26]int{}
	for _, b := range []byte(s) {
		count[b-'a']++
	}
	vowel, consonant := 0, 0
	for i := 'a'; i <= 'z'; i++ {
		if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' {
			vowel = max(vowel, count[i-'a'])
			continue
		}
		consonant = max(consonant, count[i-'a'])
	}
	return vowel + consonant
}
