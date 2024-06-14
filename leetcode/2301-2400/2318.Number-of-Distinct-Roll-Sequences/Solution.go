package Solution

const mod2318 = 1000000007

func Solution(n int) int {
	cache := make([]map[int]map[[7]int]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = map[int]map[[7]int]int{}
		for j := 1; j < 7; j++ {
			cache[i][j] = map[[7]int]int{}
		}
	}
	ref := map[int][]int{
		1: {1, 2, 3, 4, 5, 6},
		2: {1, 3, 5},
		3: {1, 2, 4, 5},
		4: {1, 3, 5},
		5: {1, 2, 3, 4, 6},
		6: {1, 5},
	}

	var dfs func(int, int, [7]int) int
	dfs = func(nn, selected int, used [7]int) int {
		if nn <= 1 {
			return nn
		}
		m := cache[nn]
		if v, ok := m[selected]; ok {
			if v1, ok1 := v[used]; ok1 {
				return v1
			}
		}
		ans := 0
		for _, next := range ref[selected] {
			x := used
			if used[next] == 0 {
				used[next] = 2
				for j := 1; j < 7; j++ {
					if j == next || used[j] == 0 {
						continue
					}
					used[j]--
				}
				ans = (ans + dfs(nn-1, next, used)) % mod2318
			}
			used = x
		}
		cache[nn][selected][used] = ans
		return ans
	}
	r := 0
	for i := 1; i < 7; i++ {
		// 1:2
		used := [7]int{}
		used[i] = 2
		r = (r + dfs(n, i, used)) % mod2318
	}
	return r
}
