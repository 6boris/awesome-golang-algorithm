package Solution

func Solution(stones []int) bool {
	cache := make(map[int]map[int]bool)
	stonePos := make(map[int]int)
	for idx, pos := range stones {
		stonePos[pos] = idx
		cache[idx] = make(map[int]bool)
	}

	var dfs func(int, int) bool
	dfs = func(index, k int) bool {
		if index == len(stones)-1 {
			// 跳上去了
			return true
		}

		if v, ok := cache[index][k]; ok {
			return v
		}

		cache[index][k] = false
		for _, step := range []int{k - 1, k, k + 1} {
			if step <= 0 {
				continue
			}
			next := stones[index] + step
			if nextIndex, ok := stonePos[next]; ok {
				if dfs(nextIndex, step) {
					cache[index][k] = true
					return true
				}
			}
		}
		return false
	}

	return dfs(0, 0)
}
