package Solution

func Solution(words []string) string {
	for _, str := range words {
		s, e := 0, len(str)-1
		for ; s < e && str[s] == str[e]; s, e = s+1, e-1 {
		}
		if s >= e {
			return str
		}
	}
	return ""
}
