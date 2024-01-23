package Solution

func uniquePattern(str string) []int {
	cache := [26]int{}
	i := 1
	cache[str[0]-'a'] = 1
	res := []int{1}
	for idx := 1; idx < len(str); idx++ {
		if r := cache[str[idx]-'a']; r != 0 {
			res = append(res, r)
			continue
		}
		i++
		res = append(res, i)
		cache[str[idx]-'a'] = i
	}
	return res
}

func Solution(words []string, pattern string) []string {
	p := uniquePattern(pattern)
	res := make([]string, 0)
	for _, word := range words {
		p1 := uniquePattern(word)
		if len(p) != len(p1) {
			continue
		}
		i := 0
		for ; i < len(p); i++ {
			if p[i] != p1[i] {
				break
			}
		}
		if i == len(p) {
			res = append(res, word)
		}
	}
	return res
}
