package Solution

const mod1155 = 1000000007

func Solution(n int, k int, target int) int {
	cache := make(map[int]map[int]int)
	for i := 1; i <= k; i++ {
		cache[i] = make(map[int]int)
	}

	var dfs func(int, int) int
	dfs = func(nn, tt int) int {
		if nn == 0 {
			if tt == 0 {
				return 1
			}
			return 0
		}
		if v, ok := cache[nn]; ok {
			if c, ok := v[tt]; ok {
				return c
			}
		}
		c := 0

		for i := 1; i <= k && i <= tt; i++ {
			c = (c + dfs(nn-1, tt-i)) % mod1155
		}
		cache[nn][tt] = c
		return c
	}
	return dfs(n, target)
}
