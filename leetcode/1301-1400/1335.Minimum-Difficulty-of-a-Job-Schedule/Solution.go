package Solution

func Solution(jobDifficulty []int, d int) int {
	l := len(jobDifficulty)
	maxCache := make([][]int, l)
	cache := make(map[int]map[int]int)

	for i := 0; i < l; i++ {
		maxCache[i] = make([]int, l)
		maxCache[i][i] = jobDifficulty[i]
		cache[i] = make(map[int]int)
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			maxCache[i][j] = maxCache[i][j-1]
			if maxCache[i][j] < jobDifficulty[j] {
				maxCache[i][j] = jobDifficulty[j]
			}
		}
	}

	var dfs func(int, int) int
	dfs = func(index, day int) int {
		if index >= l {
			return -1
		}
		if l-index < day {
			return -1
		}
		if c, ok := cache[index]; ok {
			if cnt, ok := c[day]; ok {
				return cnt
			}
		}
		if day == 1 {
			return maxCache[index][l-1]
		}
		ans := -1
		for i := index; i < l; i++ {
			cost := maxCache[index][i]
			if r := dfs(i+1, day-1); r != -1 && (ans == -1 || ans > r+cost) {
				ans = r + cost
			}
		}
		cache[index][day] = ans
		return ans
	}
	return dfs(0, d)
}
