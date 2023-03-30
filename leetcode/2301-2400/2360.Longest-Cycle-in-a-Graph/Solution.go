package Solution

func dfs2360(edges []int, start1, start, path int, hasCycle map[int]int, ans *int) {
	to := edges[start]
	if to == -1 {
		// 没有路可以走
		return
	}
	if to < 0 {
		v, ok := hasCycle[start]
		if ok && v == start1 {
			if diff := to - path; diff > *ans {
				*ans = diff
			}
		}
		return
	}
	edges[start] = path
	hasCycle[start] = start1
	dfs2360(edges, start1, to, path-1, hasCycle, ans)
}

// dfs
func Solution(edges []int) int {
	ans := -1
	hasCycle := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		if _, ok := hasCycle[i]; ok || edges[i] < -1 {
			continue
		}
		dfs2360(edges, i, i, -2, hasCycle, &ans)
	}
	return ans
}
