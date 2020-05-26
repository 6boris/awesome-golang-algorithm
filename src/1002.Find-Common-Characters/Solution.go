package Solution

import "math"

func Solution(A []string) []string {
	commonChars := make([]string, 0)
	minFrequencies := [26]int{}
	for i := range minFrequencies {
		minFrequencies[i] = math.MaxInt32
	}
	for _, word := range A {
		charFrequencies := [26]int{}
		for _, char := range word {
			charFrequencies[char-'a']++
		}
		for i := 0; i < 26; i++ {
			minFrequencies[i] = min(minFrequencies[i], charFrequencies[i])
		}
	}
	for i := 0; i < 26; i++ {
		for minFrequencies[i] > 0 {
			commonChars = append(commonChars, string(i+'a'))
			minFrequencies[i]--
		}
	}
	return commonChars
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
