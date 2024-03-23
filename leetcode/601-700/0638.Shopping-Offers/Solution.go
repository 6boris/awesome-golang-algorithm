package Solution

func Solution(price []int, special [][]int, needs []int) int {
	// dfs + cache
	l := len(price)
	cache := make(map[[6]int]int)
	cache[[6]int{0, 0, 0, 0, 0, 0}] = 0
	var dup func([6]int) [6]int
	dup = func(cur [6]int) [6]int {
		r := [6]int{}
		for i := 0; i < 6; i++ {
			r[i] = cur[i]
		}
		return r
	}
	var dfs func([6]int) int
	dfs = func(cur [6]int) int {
		if v, ok := cache[cur]; ok {
			return v
		}
		ans := 0
		for i := 0; i < l; i++ {
			ans += price[i] * cur[i]
		}

		for _, offer := range special {
			tmp := dup(cur)

			j := 0
			for ; j < l; j++ {
				if offer[j] > tmp[j] {
					break
				}
				tmp[j] -= offer[j]
			}
			if j == l {
				ans = min(ans, dfs(tmp)+offer[l])
			}
		}
		cache[cur] = ans
		return ans
	}
	x := [6]int{}
	for i := 0; i < len(needs); i++ {
		x[i] = needs[i]
	}
	return dfs(x)
}
