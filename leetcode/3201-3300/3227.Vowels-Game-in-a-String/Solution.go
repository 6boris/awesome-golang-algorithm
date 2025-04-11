package Solution

func Solution(s string) bool {
	for _, e := range s {
		if e == 'a' || e == 'e' || e == 'i' || e == 'o' || e == 'u' {
			return true
		}
	}
	return false
}
