package Solution

func power(base, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= base
	}
	return result
}

const mod = 1000000007

func Solution(n int, x int) int {
	candidates := []int{}
	i := 1
	for ; ; i++ {
		p := power(i, x)
		if p > n {
			break
		}
		candidates = append(candidates, p)
	}

	var dfs func(index, target int) int
	cache := make(map[[2]int]int)

	dfs = func(index, target int) int {
		if target == 0 {
			return 1
		}
		if target < 0 || index == len(candidates) {
			return 0
		}
		key := [2]int{index, target}
		if val, ok := cache[key]; ok {
			return val
		}
		res := dfs(index+1, target-candidates[index]) % mod
		res = (res + dfs(index+1, target)) % mod
		cache[key] = res
		return res
	}

	return dfs(0, n)
}
