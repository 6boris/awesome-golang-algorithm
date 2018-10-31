package Solution

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return make([]int, 0)
	}

	wl := len(words[0])
	ans := make([]int, 0)

	wordcount := len(words)
	wo := make(map[string]int)

	for _, word := range words {
		wo[word]++
	}
	for i := 0; i <= len(s)-wl*len(words); i++ {
		if wo[s[i:i+wl]] == 0 {
			continue
		}
		tmp := make(map[string]int)
		//found a word in dict
		count := 0

		for j := i; j <= len(s)-wl && count < wordcount; j += wl {
			word := s[j : j+wl]
			tmp[word]++

			if tmp[word] > wo[word] {
				break
			}
			count++
		}
		if count == wordcount {
			ans = append(ans, i)
		}
	}
	return ans
}
