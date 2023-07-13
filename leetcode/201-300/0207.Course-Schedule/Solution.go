package Solution

func Solution(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		// 随便学
		return true
	}

	order := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, item := range prerequisites {
		who, dep := item[0], item[1]
		inDegree[who]++
		if _, ok := order[dep]; !ok {
			order[dep] = make([]int, 0)
		}
		order[dep] = append(order[dep], who)
	}

	queue := make([]int, 0)
	left := numCourses
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
			left--
		}
	}

	for len(queue) > 0 {
		nextQ := make([]int, 0)
		for _, item := range queue {
			for _, next := range order[item] {
				inDegree[next]--
				if inDegree[next] == 0 {
					nextQ = append(nextQ, next)
				}
			}
		}
		left -= len(nextQ)
		if len(nextQ) == 0 {
			break
		}
		queue = nextQ
	}
	return left == 0
}
