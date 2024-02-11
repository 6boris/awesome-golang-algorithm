package Solution

func Solution(words []string) bool {
	letters := [26]int{}
	for _, word := range words {
		for _, b := range word {
			letters[b-'a']++
		}
	}
	l := len(words)
	for i := 0; i < 26; i++ {
		if letters[i]%l != 0 {
			return false
		}
	}
	return true
}
