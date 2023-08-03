package Solution

func Solution(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	a := make([]int, 26)
	b := make([]int, 26)
	diff := 0
	//aa, aa
	for i := 0; i < len(s); i++ {
		if s[i] == goal[i] {
			a[s[i]-'a']++
			continue
		}
		diff++
		b[goal[i]-'a']++
		b[s[i]-'a']--
	}
	// ab,ab
	if diff == 0 {
		// ab,ab
		// aa, aa
		// diff 是0 的时候表明每个位置，都是一样的，需要看是否有重复元素，有就可以调整
		for i := 0; i < 26; i++ {
			if a[i] >= 2 {
				return true
			}
		}
		return false
	}
	// aac, bac
	if diff == 2 {
		for i := 0; i < 26; i++ {
			if b[i] != 0 {
				return false
			}
		}
		return true
	}
	return false
}
