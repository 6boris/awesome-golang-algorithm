package Solution

func Solution(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}

	mapping := make(map[byte]byte)
	reverseMapping := make(map[byte]struct{})
	for idx := 0; idx < ls; idx++ {
		if v, ok := mapping[s[idx]]; ok {
			if v != t[idx] {
				return false
			}
			continue
		}
		if _, ok := reverseMapping[t[idx]]; ok {
			return false
		}

		mapping[s[idx]] = t[idx]
		reverseMapping[t[idx]] = struct{}{}
	}

	return true
}
