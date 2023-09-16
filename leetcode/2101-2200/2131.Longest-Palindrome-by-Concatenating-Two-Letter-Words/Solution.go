package Solution

func revWord(s string) string {
	bs := []byte(s)
	for s, e := 0, len(s)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}
	return string(bs)
}

func isAllByteSame(s string) bool {
	cmp := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != cmp {
			return false
		}
	}
	return true
}

func Solution(words []string) int {
	diffStore := make(map[string]int)
	sameStore := make(map[string]int)

	for _, word := range words {
		if isAllByteSame(word) {
			sameStore[word]++
			continue
		}
		diffStore[word]++
	}

	samMax, diffMax := 0, 0
	evenSum := 0
	for k, v := range sameStore {
		if v&1 == 0 {
			evenSum += v * len(k)
			continue
		}
		evenSum += (v - 1) * len(k)
		if samMax == 0 || len(k) > samMax {
			samMax = len(k)
		}
	}
	samMax += evenSum

	for _, word := range words {
		if _, ok := sameStore[word]; ok {
			continue
		}
		rev := revWord(word)
		c1, ok1 := diffStore[word]
		c2, ok2 := diffStore[rev]
		if ok1 && ok2 {
			s := c1
			if c2 < s {
				s = c2
			}
			diffMax += s * len(word) * 2
			diffStore[word] -= s
			diffStore[rev] -= s
		}
	}
	return diffMax + samMax
}
