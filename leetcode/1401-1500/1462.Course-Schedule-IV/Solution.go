package Solution

func Solution(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	father := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		father[i] = make([]bool, numCourses)
	}

	in := make([]int, numCourses)
	release := make(map[int][]int)
	for _, require := range prerequisites {
		in[require[1]]++
		if _, ok := release[require[0]]; !ok {
			release[require[0]] = make([]int, 0)
		}
		release[require[0]] = append(release[require[0]], require[1])
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		nq := make([]int, 0)
		for _, cur := range queue {
			for _, next := range release[cur] {
				in[next]--
				father[next][cur] = true
				for k, v := range father[cur] {
					if v {
						father[next][k] = true
					}
				}
				if in[next] == 0 {
					nq = append(nq, next)
				}
			}
		}
		queue = nq
	}
	ans := make([]bool, len(queries))
	for idx, query := range queries {
		ans[idx] = father[query[1]][query[0]]
	}
	return ans
}
