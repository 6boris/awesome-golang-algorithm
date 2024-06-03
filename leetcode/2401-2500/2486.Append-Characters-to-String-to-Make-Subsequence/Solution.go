package Solution

func Solution(s string, t string) int {
	match := 0
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			match++
			i++
			j++
			continue
		}
		i++
	}
	return len(t) - match
}
