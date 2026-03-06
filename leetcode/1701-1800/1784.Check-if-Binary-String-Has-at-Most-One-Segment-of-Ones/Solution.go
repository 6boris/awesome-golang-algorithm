package Solution

func Solution(s string) bool {
	seg := 0
	prev := s[0]
	if prev == '1' {
		seg++
	}

	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			continue
		}
		prev = s[i]
		if prev == '1' {
			seg++
		}
		if seg > 1 {
			return false
		}
	}
	return true
}
