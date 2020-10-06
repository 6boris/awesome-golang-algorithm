package Solution

func firstUniqChar(s string) byte {
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}

	for _, v := range s {
		if m[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}
