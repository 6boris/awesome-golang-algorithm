package Solution

func Solution(s string) int {
	var ret int
	cnt := [26]int{}
	start, end := 0, 0
	for ; end < len(s); end++ {
		i := s[end] - 'a'
		if cnt[i] < 2 {
			cnt[i]++
			continue
		}
		ret = max(ret, end-start)
		for ; s[start] != s[end]; start++ {
			cnt[s[start]-'a']--
		}
		start++
	}
	ret = max(ret, end-start)

	return ret
}
