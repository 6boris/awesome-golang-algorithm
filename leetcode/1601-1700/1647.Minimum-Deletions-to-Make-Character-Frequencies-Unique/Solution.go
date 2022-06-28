package Solution

func Solution(s string) int {
	letters := make([]int, 26)
	for _, b := range []byte(s) {
		letters[b-'a']++
	}
	occurrences := make(map[int]struct{})
	ans := 0
	for idx := 0; idx < 26; idx++ {
		if letters[idx] == 0 {
			continue
		}
		r := letters[idx]
		if _, ok := occurrences[r]; !ok {
			occurrences[r] = struct{}{}
			continue
		}
		for r = r - 1; r >= 0; r-- {
			ans++
			if _, ok := occurrences[r]; !ok {
				occurrences[r] = struct{}{}
				break
			}
		}
	}

	return ans
}
