package Solution

func Solution(s1 string, s2 string) bool {
	a := make([]int, 26)
	cache := make(map[string]bool)
	return isScramble87([]byte(s1), []byte(s2), cache, chars1(a))
}

func chars1(baseS1 []int) func([]byte, []byte) bool {
	return func(s1, s2 []byte) bool {
		for i := 0; i < 26; i++ {
			baseS1[i] = 0
		}
		for i := 0; i < len(s1); i++ {
			baseS1[s1[i]-'a']++
			baseS1[s2[i]-'a']--
		}
		for i := 0; i < 26; i++ {
			if baseS1[i] != 0 {
				return false
			}
		}
		return true
	}
}

func isScramble87(s1, s2 []byte, cache map[string]bool, cmp func([]byte, []byte) bool) bool {
	if v, ok := cache[string(s1)+string(s2)]; ok {
		return v
	}
	if string(s1) == string(s2) {
		return true
	}
	ll := len(s1)
	equal := true
	for i := 0; i < ll; i++ {
		if s1[i] != s2[ll-1-i] {
			equal = false
			break
		}
	}
	if equal {
		return true
	}
	if !cmp(s1, s2) {
		return false
	}
	cache[string(s1)+string(s2)] = false
	for l := 1; l < ll; l++ {
		if isScramble87(s1[:l], s2[:l], cache, cmp) && isScramble87(s1[l:], s2[l:], cache, cmp) {
			cache[string(s1)+string(s2)] = true
			return true
		}
		// a,b/c,    cba
		if isScramble87(s1[:l], s2[ll-l:], cache, cmp) && isScramble87(s1[l:], s2[:ll-l], cache, cmp) {
			cache[string(s1)+string(s2)] = true
			return true
		}
	}
	return false
}
