package Solution

func Solution(allowed string, words []string) int {
	in := [26]bool{}
	for _, b := range allowed {
		in[b-'a'] = true
	}
	ans := 0
	for _, word := range words {
		ok := true
		for _, b := range word {
			if !in[b-'a'] {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}
