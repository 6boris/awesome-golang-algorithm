package Solution

func Solution(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return 0
	}
	// bfs
	step := 0
	indices := make(map[int][]int)
	visited := make([]bool, length)
	for i := 0; i < length; i++ {
		if _, ok := indices[arr[i]]; !ok {
			indices[arr[i]] = make([]int, 0)
		}
		indices[arr[i]] = append(indices[arr[i]], i)
	}
	visited[0] = true
	queue := []int{0}

	for len(queue) > 0 {
		nextQ := make([]int, 0)
		for _, index := range queue {
			if index-1 >= 0 && !visited[index-1] {
				visited[index-1] = true
				nextQ = append(nextQ, index-1)
			}
			if index+1 < length && !visited[index+1] {
				visited[index+1] = true
				nextQ = append(nextQ, index+1)
			}
			for _, rel := range indices[arr[index]] {
				if !visited[rel] {
					visited[rel] = true
					nextQ = append(nextQ, rel)
				}
			}
			// 注意这里的超时点
			delete(indices, arr[index])
		}
		step++
		if visited[length-1] {
			return step
		}
		queue = nextQ

	}
	return 0
}
