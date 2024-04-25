package Solution

func Solution(s string, k int) int {
	dis := [26]int{}
	dis[s[0]-'a'] = 1
	for i := 1; i < len(s); i++ {
		td := uint8(0)
		for ; td <= uint8(k); td++ {
			if next := s[i] + td; next >= 'a' && next <= 'z' {
				if dis[next-'a'] != 0 && dis[next-'a']+1 > dis[s[i]-'a'] {
					dis[s[i]-'a'] = dis[next-'a'] + 1
				}
			}

		}
		td = uint8(1)
		for ; td <= uint8(k); td++ {
			if pre := s[i] - td; pre >= 'a' && pre <= 'z' {
				if dis[pre-'a'] != 0 && dis[pre-'a']+1 > dis[s[i]-'a'] {
					dis[s[i]-'a'] = dis[pre-'a'] + 1
				}
			}
		}
		if dis[s[i]-'a'] == 0 {
			dis[s[i]-'a'] = 1
		}
	}
	ans := 1
	for i := 0; i < 26; i++ {
		if dis[i] > ans {
			ans = dis[i]
		}
	}
	return ans
}
