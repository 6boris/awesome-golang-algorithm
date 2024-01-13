package Solution

func Solution(s string, t string) int {
	as := [26]int{}
	ts := [26]int{}
	if s == t {
		return 0
	}

	for _, b := range s {
		as[b-'a']++
	}
	for _, b := range t {
		ts[b-'a']++
	}

	more := 0
	for i := 0; i < 26; i++ {
		diff := ts[i] - as[i]
		if diff == 0 {
			continue
		}
		if diff > 0 {
			more += diff
		}
	}
	return more
}
