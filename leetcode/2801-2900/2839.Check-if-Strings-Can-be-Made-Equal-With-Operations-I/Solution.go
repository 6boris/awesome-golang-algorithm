package Solution

func Solution(s1 string, s2 string) bool {
	testIndex := func(i, j int) bool {
		a1, b1 := s1[i], s1[j]
		a2, b2 := s2[i], s2[j]
		if a1 < b1 {
			a1, b1 = b1, a1
		}
		if a2 < b2 {
			a2, b2 = b2, a2
		}
		return a1 == a2 && b1 == b2
	}
	return testIndex(0, 2) && testIndex(1, 3)
}
