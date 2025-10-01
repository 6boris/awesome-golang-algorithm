package Solution

func Solution(n int, preferences [][]int, pairs [][]int) int {
	likeIndex := make([]map[int]int, len(preferences))
	for i, p := range preferences {
		likeIndex[i] = make(map[int]int)
		for idx, like := range p {
			likeIndex[i][like] = idx
		}
	}

	pair := make(map[int]int)
	for _, p := range pairs {
		pair[p[0]] = p[1]
		pair[p[1]] = p[0]
	}
	var ret int
	for _, p := range pairs {
		x, y := p[0], p[1]
		for _, u := range preferences[x] {
			if u == y {
				break
			}
			if v, ok := pair[u]; ok {
				if likeIndex[u][x] < likeIndex[u][v] {
					ret++
					break
				}
			}
		}

		for _, u := range preferences[y] {
			if u == x {
				break
			}
			if v, ok := pair[u]; ok {
				if likeIndex[u][y] < likeIndex[u][v] {
					ret++
					break
				}
			}
		}
	}
	return ret
}
