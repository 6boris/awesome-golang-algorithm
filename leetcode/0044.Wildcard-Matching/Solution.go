package Solution

func isMatch(s, p string) bool {
	prev, now := make([]bool, len(p)+1), make([]bool, len(p)+1)
	for i := 0; i <= len(s); i++ {
		now, prev = prev, now
		now[0] = i == 0
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				now[j] = prev[j] || prev[j-1] || now[j-1]
			} else {
				now[j] = prev[j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
		}
	}
	return now[len(p)]
}
