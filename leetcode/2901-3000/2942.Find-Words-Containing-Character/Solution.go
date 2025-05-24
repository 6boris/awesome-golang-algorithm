package Solution

func Solution(words []string, x byte) []int {
	var ans []int
	for i, word := range words {
		ok := false
		for _, b := range []byte(word) {
			if b == x {
				ok = true
				break
			}
		}
		if ok {
			ans = append(ans, i)
		}
	}
	return ans
}
