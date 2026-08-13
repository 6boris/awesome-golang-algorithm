package Solution

func Solution(s string) int {
	e := make(map[byte]struct{})
	for i := range s {
		e[s[i]-'a'] = struct{}{}
	}
	return len(e)
}
