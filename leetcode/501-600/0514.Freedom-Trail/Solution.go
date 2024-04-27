package Solution

func Solution(ring string, key string) int {
	// 感觉是dfs+cache实现
	pos := [26][]int{}
	for i := 0; i < 26; i++ {
		pos[i] = make([]int, 0)
	}
	for i, r := range ring {
		pos[r-'a'] = append(pos[r-'a'], i)
	}
	lr := len(ring)
	lk := len(key)
	cache := make([][]int, lr)
	for i := 0; i < lr; i++ {
		cache[i] = make([]int, lk)
	}

	var dfs func(int, int) int
	dfs = func(keyIndex, ringIndex int) int {
		if keyIndex == len(key) {
			return 0
		}
		if cache[ringIndex][keyIndex] != 0 {
			return cache[ringIndex][keyIndex]
		}

		curByte := key[keyIndex]
		ans := -1
		for _, idx := range pos[curByte-'a'] {
			diff := ringIndex - idx
			if diff < 0 {
				diff = -diff
			}
			diff = min(lr-diff, diff)
			diff++
			steps := dfs(keyIndex+1, idx)
			if ans == -1 || diff+steps < ans {
				ans = diff + steps
			}
		}
		cache[ringIndex][keyIndex] = ans
		return ans
	}
	return dfs(0, 0)
}
