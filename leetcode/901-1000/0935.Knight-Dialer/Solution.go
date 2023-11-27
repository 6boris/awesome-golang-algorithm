package Solution

const mod935 = 1000000007

func Solution(n int) int {
	nexts := map[int][]int{
		1: {6, 8},
		2: {7, 9},
		3: {4, 8},
		4: {0, 3, 9},
		6: {0, 1, 7},
		7: {2, 6},
		8: {1, 3},
		9: {4, 2},
		0: {4, 6},
	}
	cache := make(map[int]map[int]int)
	for i := 0; i <= 9; i++ {
		cache[i] = make(map[int]int)
		cache[i][1] = 1
	}
	var dfs func(int, int) int
	dfs = func(num, nn int) int {
		v, ok := cache[num]
		if ok {
			vv, ok1 := v[nn]
			if ok1 {
				return vv
			}
		}
		ans := 0
		for _, next := range nexts[num] {
			ans = (ans + dfs(next, nn-1)) % mod935
		}
		cache[num][nn] = ans
		return ans
	}
	ans := 0
	for i := 0; i <= 9; i++ {
		a := dfs(i, n) % mod935
		ans = (ans + a) % mod935
	}
	return ans
}
