package Solution

func Solution(s string) int {
	ans := 0
	total := [26]int{}
	for _, b := range s {
		total[b-'a']++
	}
	left := [26]int{}
	for end := 0; end < len(s)-1; end++ {
		left[s[end]-'a']++
		a, b := 0, 0
		for i := 0; i < 26; i++ {
			if left[i] > 0 {
				a++
			}
			if total[i] > left[i] {
				b++
			}
		}
		if a == b {
			ans++
		}
	}
	return ans
}
