package Solution

func Solution(points [][]int) int {
	length := len(points)
	distance := make([][]int, length)
	visited := make([]bool, length)
	for i := 0; i < length; i++ {
		distance[i] = make([]int, length)
		for j := 0; j < length; j++ {
			distance[i][j] = -1
		}
	}

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]

			if x < 0 {
				x = -x
			}
			if y < 0 {
				y = -y
			}
			distance[i][j] = x + y
			distance[j][i] = x + y
		}
	}

	// 选一个起点
	// 第一个点被选中
	visited[0] = true
	ans := 0
	for {
		selectIndex, selectDistance := -1, -1
		for i, v := range visited {
			if !v {
				continue
			}
			// 遍历过，则尝试选择一个较小的距离
			for rel, dis := range distance[i] {
				if dis == -1 || visited[rel] {
					// 不可达
					continue
				}
				if selectDistance == -1 || selectDistance > dis {
					selectIndex = rel
					selectDistance = dis
				}
			}
		}
		if selectIndex == -1 {
			break
		}
		visited[selectIndex] = true
		ans += selectDistance
	}

	return ans
}
