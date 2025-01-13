package Solution

func Solution(s string) int {
	count := [26]int{}
	for _, b := range s {
		count[b-'a']++
	}
	ans := 0
	for i := range 26 {
		if count[i] != 0 {
			if count[i]&1 == 1 {
				ans++
				continue
			}
			ans += 2
		}
	}
	return ans
}
