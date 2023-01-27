package Solution

import "strings"

func Solution(words []string) []string {
	ans := make([]string, 0)
	for _, w := range words {
		cache := make(map[int]bool)
		if findAllConcatenatedWordsInADictAux(w, 0, words, cache) {
			ans = append(ans, w)
		}
		/*
			try(w, 0, 0, words, c)
			if c[w] >= 2 {
				ans = append(ans, w)
			}
		*/
	}
	return ans
}

/*
// g,e,ge

	func try(word string, start, length int, words []string, cache map[string]int) {
		if start == len(word) {
			cache[word] = length
			return
		}
		if v, ok := cache[word[start:]]; ok {
			if length+v > cache[word] {
				cache[word] = length+v
			}
			return
		}
		for _, w := range words {
			if word == w {
				continue
			}
			if strings.HasPrefix(word[start:], w) {
				cache[word[:start+1]] = length
				try(word, start+len(w), length+1, words, cache)
			}
		}
	}
*/
func findAllConcatenatedWordsInADictAux(s string, idx int, words []string, check map[int]bool) bool {
	if idx == len(s) {
		check[idx] = true
		return true
	}
	if val, ok := check[idx]; ok {
		return val
	}
	for _, word := range words {
		if word == s {
			continue
		}
		if strings.HasPrefix(s[idx:], word) && findAllConcatenatedWordsInADictAux(s, idx+len(word), words, check) {
			check[idx] = true
			return true
		}
	}
	check[idx] = false
	return false
}
