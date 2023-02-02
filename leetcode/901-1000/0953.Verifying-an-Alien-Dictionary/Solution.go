package Solution

func Solution(words []string, order string) bool {
	newOrder := make([]uint8, 26)
	for i := 0; i < 26; i++ {
		newOrder[order[i]-'a'] = uint8(i)
	}

	var less func(string, string) bool
	less = func(a, b string) bool {
		la, lb := len(a), len(b)
		for i := 0; i < la && i < lb; i++ {
			if newOrder[a[i]-'a'] == newOrder[b[i]-'a'] {
				continue
			}
			return newOrder[a[i]-'a'] < newOrder[b[i]-'a']
		}
		return la <= lb
	}
	for i := 1; i < len(words); i++ {
		if !less(words[i-1], words[i]) {
			return false
		}
	}
	return true
}
