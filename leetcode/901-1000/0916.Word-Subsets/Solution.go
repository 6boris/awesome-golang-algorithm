package Solution

func Solution(words1 []string, words2 []string) []string {
	ac := [26]int{}
	for _, w := range words2 {
		tmp := [26]int{}
		for _, b := range w {
			tmp[b-'a']++
		}
		for i := range 26 {
			if tmp[i] > ac[i] {
				ac[i] = tmp[i]
			}
		}
	}
	ans := make([]string, 0)
	for _, a := range words1 {
		tmp := [26]int{}
		for _, b := range a {
			tmp[b-'a']++
		}
		ok := true
		for i := range 26 {
			if ac[i] > tmp[i] {
				ok = false
				break
			}
		}
		if ok {
			ans = append(ans, a)
		}
	}
	return ans
}
