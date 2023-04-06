package Solution

func Solution(s string) int {
	chars := make([]bool, 26)
	ans := 1
	for _, b := range s {
		if !chars[b-'a'] {
			chars[b-'a'] = true
			continue
		}
		ans++
		for i := 0; i < 26; i++ {
			chars[i] = false
		}
		chars[b-'a'] = true
	}
	return ans
}
