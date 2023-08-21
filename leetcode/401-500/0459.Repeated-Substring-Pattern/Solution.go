package Solution

func Solution(s string) bool {
	length := len(s)
	for l := length - 1; l >= 1; l-- {
		if length%l != 0 {
			continue
		}
		base := s[:l]
		idx := l
		for ; idx < length; idx += l {
			if base != s[idx:idx+l] {
				break
			}
		}
		if idx >= length {
			return true
		}
	}
	return false
}
