package Solution

func next1967(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	if n == 0 {
		return next
	}
	j := 0
	next[0] = 0
	for i := 1; i < n; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func Solution(patterns []string, word string) int {
	var (
		ret       int
		kmpSearch func(string) int
	)
	kmpSearch = func(pattern string) int {
		next := next1967(pattern)
		m := len(word)
		j := 0
		for i := 0; i < m; i++ {
			for j > 0 && word[i] != pattern[j] {
				j = next[j-1]
			}
			if word[i] == pattern[j] {
				j++
			}
			if j == len(pattern) {
				return i - len(pattern) + 1

			}
		}
		return -1
	}
	for i := range patterns {
		if idx := kmpSearch(patterns[i]); idx != -1 {
			ret++
		}
	}
	return ret
}
