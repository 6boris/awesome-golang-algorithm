package Solution

func Solution(s string) int {
	ans := 0
	cur := s[0]
	c := 1
	for i := 1; i < len(s); i++ {
		if s[i] == cur {
			c = 1 - c
			continue
		}
		if c == 0 {
			c = 1
			cur = s[i]
			continue
		}
		ans++
		i++
		if i != len(s) {
			cur = s[i]
			c = 1
		}
	}
	return ans
}
