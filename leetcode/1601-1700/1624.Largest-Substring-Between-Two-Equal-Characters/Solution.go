package Solution

func Solution(s string) int {
	m1 := [26]int{}
	m2 := [26]int{}
	for i := 0; i < 26; i++ {
		m1[i] = -1
		m2[i] = -1
	}
	for idx, b := range s {
		if m1[b-'a'] == -1 {
			m1[b-'a'] = idx
		}
	}
	for idx := len(s) - 1; idx >= 0; idx-- {
		if m2[s[idx]-'a'] == -1 {
			m2[s[idx]-'a'] = idx
		}
	}
	ans := -1
	for i := 0; i < 26; i++ {
		if m1[i] != -1 && m2[i] != -1 && m2[i] != m1[i] {
			if diff := m2[i] - m1[i] - 1; diff > ans {
				ans = diff
			}
		}
	}
	return ans
}
