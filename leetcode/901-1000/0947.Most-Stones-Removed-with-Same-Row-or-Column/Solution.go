package Solution

func Solution(stones [][]int) int {
	visited := make([]bool, len(stones))
	combindStones := 0
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for idx := 0; idx < len(stones); idx++ {
			if visited[idx] {
				continue
			}

			// 横着或者竖着
			if canRemove(stones[i], stones[idx]) {
				combindStones++
				dfs(idx)
			}
		}
	}
	for i := 0; i < len(stones); i++ {
		if visited[i] {
			continue
		}
		dfs(i)
	}

	return combindStones
}

func canRemove(p1, p2 []int) bool {
	return p1[0] == p2[0] || p1[1] == p2[1]
}
