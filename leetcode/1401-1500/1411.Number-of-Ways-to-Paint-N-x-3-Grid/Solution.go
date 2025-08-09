package Solution

const mod = 1000000007

func Solution(n int) int {
	rel := [12][]int{
		{1, 4, 10, 2, 5},
		{0, 3, 6, 8, 11},
		{0, 3, 6, 7},

		{1, 7, 10, 2},
		{0, 6, 9, 8},
		{0, 6, 9, 7, 10},

		{1, 4, 2, 5, 11},
		{3, 2, 5, 11},
		{9, 1, 4, 10},

		{4, 5, 8, 11},
		{0, 3, 5, 8, 11},
		{1, 7, 10, 6, 9},
	}
	var dfs func(int, int) int
	cache := make(map[[2]int]int)
	dfs = func(parent, left int) int {
		if left == 0 {
			return 1
		}
		ans := 0
		l := len(rel[parent])
		key := [2]int{l, left}
		if v, ok := cache[key]; ok {
			return v
		}
		for _, next := range rel[parent] {
			ans = (ans + dfs(next, left-1)) % mod
		}
		cache[key] = ans
		return ans
	}
	a := (dfs(0, n-1) * 6) % mod
	b := (dfs(2, n-1) * 6) % mod
	return (a + b) % mod
}
