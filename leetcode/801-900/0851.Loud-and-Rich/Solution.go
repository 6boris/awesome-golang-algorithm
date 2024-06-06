package Solution

func Solution(richer [][]int, quiet []int) []int {
	l := len(quiet)
	ans := make([]int, l)
	greater := make([][]bool, l)
	for i := 0; i < l; i++ {
		greater[i] = make([]bool, l)
	}

	for _, cmp := range richer {
		greater[cmp[1]][cmp[0]] = true
	}

	cache := make(map[int]int)
	var dfs func(int) int
	dfs = func(x int) int {
		if v, ok := cache[x]; ok {
			return v
		}
		ans := x
		for u, v := range greater[x] {
			if v {
				if quiet[u] < quiet[ans] {
					ans = u
				}
				if r := dfs(u); quiet[ans] > quiet[r] {
					ans = r
				}
			}
		}
		cache[x] = ans
		return ans
	}
	for i := 0; i < l; i++ {
		ans[i] = dfs(i)
	}
	return ans
}
