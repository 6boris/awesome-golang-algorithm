package Solution

func Solution(words []string, chars string) int {
	charFreq := [26]int{}
	for _, char := range chars {
		charFreq[char-97]++
	}
	res := 0
	for _, word := range words {
		if checkIfGood(word, charFreq) {
			res += len(word)
		}
	}
	return res
}

func checkIfGood(word string, charFreq [26]int) bool {
	for _, char := range word {
		charFreq[char-97]--
		if charFreq[char-97] < 0 {
			return false
		}
	}
	return true
}
