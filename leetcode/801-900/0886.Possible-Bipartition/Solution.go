package Solution

func Solution(n int, dislikes [][]int) bool {
	dislike := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dislike[i] = make([]bool, n+1)
	}
	for _, pair := range dislikes {
		dislike[pair[0]][pair[1]] = true
		dislike[pair[1]][pair[0]] = true
	}

	var dfs func(int, int) bool
	color := make([]int, n+1)
	dfs = func(i, nodeColor int) bool {
		color[i] = nodeColor
		for rel := 1; rel < len(dislike[i]); rel++ {
			if dislike[i][rel] && color[rel] == nodeColor {
				return false
			}
			if dislike[i][rel] && color[rel] == 0 {
				if !dfs(rel, 3-nodeColor) {
					return false
				}
			}
		}
		return true
	}

	for p := 1; p <= n; p++ {
		if color[p] == 0 {
			if !dfs(p, 1) {
				return false
			}
		}
	}
	return true
}
