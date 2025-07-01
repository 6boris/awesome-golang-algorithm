package Solution

func Solution(word string) int {
	ans := 1
	c := 0
	pre := byte(' ')
	for _, b := range []byte(word) {
		if b == pre {
			c++
			continue
		}
		if c > 1 {
			ans += c - 1
		}
		c = 1
		pre = b
	}
	if c > 1 {
		ans += c - 1
	}
	return ans
}
