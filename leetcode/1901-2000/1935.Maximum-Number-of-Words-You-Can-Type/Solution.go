package Solution

import "strings"

func Solution(text string, brokenLetters string) int {
	broken := [26]bool{}
	for _, b := range brokenLetters {
		broken[b-'a'] = true
	}
	ret, i := 0, 0
	for _, word := range strings.Split(text, " ") {
		for i = 0; i < len(word); i++ {
			if broken[word[i]-'a'] {
				break
			}
		}
		if i == len(word) {
			ret++
		}
	}
	return ret
}
