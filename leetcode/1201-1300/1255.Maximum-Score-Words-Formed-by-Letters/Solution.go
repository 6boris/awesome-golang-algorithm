package Solution

func Solution(words []string, letters []byte, score []int) int {
	var (
		match  func(a, b [26]int) bool
		search func(int, int, [26]int)
	)

	letterCount := [26]int{}
	for _, c := range letters {
		letterCount[c-'a']++
	}
	match = func(a, b [26]int) bool {
		for i := 0; i < 26; i++ {
			if a[i] > b[i] {
				return false
			}
		}
		return true
	}

	restoreWords := make([]string, 0)
	wordHash := make(map[string][26]int)
	for _, word := range words {
		store := [26]int{}
		for _, b := range word {
			store[b-'a']++
		}

		if match(store, letterCount) {
			restoreWords = append(restoreWords, word)
			wordHash[word] = store
		}
	}

	ans := 0
	search = func(index int, s int, count [26]int) {
		if index == len(restoreWords) {
			if s > ans {
				ans = s
			}
			return
		}

		w := restoreWords[index]
		if match(wordHash[w], count) {
			tmp := s
			for i := 0; i < 26; i++ {
				count[i] -= wordHash[w][i]
				tmp += wordHash[w][i] * score[i]
			}
			search(index+1, tmp, count)

			for i := 0; i < 26; i++ {
				count[i] += wordHash[w][i]
			}
		}
		search(index+1, s, count)
	}

	search(0, 0, letterCount)

	return ans
}
