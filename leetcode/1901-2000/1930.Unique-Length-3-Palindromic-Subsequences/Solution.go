package Solution

func Solution(s string) int {
	firstIndex := [26]int{}
	lastIndex := [26]int{}
	chars := make([][26]int, len(s))
	for i := 0; i < 26; i++ {
		firstIndex[i] = -1
		lastIndex[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if firstIndex[s[i]-'a'] == -1 {
			firstIndex[s[i]-'a'] = i
		}
		lastIndex[s[i]-'a'] = i
		chars[i][s[i]-'a']++
		if i != 0 {
			for idx := 0; idx < 26; idx++ {
				chars[i][idx] += chars[i-1][idx]
			}
		}
	}

	ans := 0
	for i := 0; i < 26; i++ {
		first := firstIndex[i]
		last := lastIndex[i]
		if first == last {
			continue
		}
		byteCount := 0
		for i := 0; i < 26; i++ {
			if r := chars[last-1][i] - chars[first][i]; r > 0 {
				byteCount++
			}
		}
		ans += byteCount
	}
	return ans
}
