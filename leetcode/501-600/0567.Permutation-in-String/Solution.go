package Solution

func Solution(s1 string, s2 string) bool {
	l2 := len(s2)
	l1 := len(s1)
	if l2 < l1 {
		return false
	}
	byteCount := make([]int, 26)
	for _, b := range s1 {
		byteCount[b-'a']++
	}
	cmp := make([]int, 26)
	for i := 0; i < l1; i++ {
		cmp[s2[i]-'a']++
	}
	var equal func([]int, []int) bool
	equal = func(a, b []int) bool {
		for i := 0; i < 26; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	if equal(byteCount, cmp) {
		return true
	}

	for next := l1; next < l2; next++ {
		del := next - l1
		cmp[s2[next]-'a']++
		cmp[s2[del]-'a']--
		if equal(cmp, byteCount) {
			return true
		}
	}
	return false
}
