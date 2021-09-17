package Solution

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}
	return -1
}
