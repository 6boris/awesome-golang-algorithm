package Solution

func findRepeatedDnaSequences(s string) []string {
	ans := make([]string, 0)
	if len(s) < 10 {
		return ans
	}

	cache := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		curr := s[i : i+10]
		if cache[curr] == 1 {
			ans = append(ans, curr)
		}
		cache[curr] += 1
	}
	return ans
}
