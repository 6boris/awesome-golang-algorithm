package Solution

func Solution(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	var a, b uint8
	diff := 0
	for i := range s1 {
		if s1[i] == s2[i] {
			continue
		}
		diff++
		if diff == 1 {
			a, b = s1[i], s2[i]
		} else if diff == 2 {
			// a, s1, b s2
			if !(a == s2[i] && b == s1[i]) {
				return false
			}
		} else {
			return false
		}
	}
	return diff != 1
}
