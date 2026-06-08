package Solution

func distance(x1, y1, x2, y2 int) int64 {
	diffX := int64(x2 - x1)
	diffY := int64(y2 - y1)
	return diffX*diffX + diffY*diffY
}

// 如果擦边就算，并查集可解决
func Solution(bombs [][]int) int {
	length := len(bombs)
	points := make(map[int]map[int]struct{}, length)
	for i := 0; i < length; i++ {
		points[i] = make(map[int]struct{})
	}
	for i := 0; i < len(bombs); i++ {
		ii := bombs[i]
		for j := 0; j < len(bombs); j++ {
			if i == j {
				continue
			}
			jj := bombs[j]
			dis := distance(ii[0], ii[1], jj[0], jj[1])
			actual := int64(ii[2] * ii[2])
			if actual >= dis {
				points[i][j] = struct{}{}
			}
		}
	}

	var dfs func(int, map[int]struct{}) int
	dfs = func(start int, visited map[int]struct{}) int {
		visited[start] = struct{}{}
		total := 1 // 自己
		for next := range points[start] {
			if _, ok := visited[next]; !ok {
				visited[next] = struct{}{}
				total += dfs(next, visited)
			}
		}
		return total
	}
	ans := 0
	for i := 0; i < length; i++ {
		if a := dfs(i, map[int]struct{}{}); a > ans {
			ans = a
		}
	}
	return ans
}
