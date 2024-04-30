package Solution

func Solution(s string) int {
	state := map[int]int{
		0: -1,
	}

	mask := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		shift := s[i] - '0'
		mask ^= (1 << shift)
		if v, ok := state[mask]; ok {
			ans = max(ans, i-v)
		} else {
			state[mask] = i
		}
		for j := 0; j < 10; j++ {
			if v, ok := state[mask^(1<<j)]; ok {
				ans = max(ans, i-v)
			}
		}
	}
	return ans
}
