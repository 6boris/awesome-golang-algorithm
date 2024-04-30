package Solution

func Solution(word string) int64 {
	// 10bit
	ans := int64(0)
	l := len(word)
	state := make(map[int]int64)
	state[0] = 1
	mask := 0
	for i := 0; i < l; i++ {
		mask ^= (1 << int(word[i]-'a'))
		c := state[mask]
		ans += c
		state[mask] = c + 1
		for i := 0; i < 10; i++ {
			ans += state[mask^(1<<i)]
		}
	}
	return ans
}
