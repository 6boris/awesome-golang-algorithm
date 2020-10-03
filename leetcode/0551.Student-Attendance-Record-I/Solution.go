package Solution

func checkRecord(s string) bool {
	a, l := 0, 0

	for i, _ := range s {
		if s[i] == 'A' {
			a++
		}
		if s[i] == 'L' {
			l++
		} else {
			l = 0
		}

		if a > 1 || l > 2 {
			return false
		}

	}

	return true
}
