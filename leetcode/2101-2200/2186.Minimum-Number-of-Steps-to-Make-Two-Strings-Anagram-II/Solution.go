package Solution

func Solution(s string, t string) int {
	ss, tt := [26]int{}, [26]int{}
	bs, bt := []byte(s), []byte(t)

	for i := range bs {
		ss[bs[i]-'a']++
	}
	for i := range bt {
		tt[bt[i]-'a']++
	}

	var ret, diff int
	for i := range 26 {
		if ss[i] == tt[i] {
			continue
		}
		diff = ss[i] - tt[i]
		if diff < 0 {
			diff = -diff
		}
		ret += diff
	}
	return ret
}
