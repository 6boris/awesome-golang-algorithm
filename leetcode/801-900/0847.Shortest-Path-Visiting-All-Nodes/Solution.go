package Solution

func Solution(graph [][]int) int {
	queue := [][]int{}
	size := len(graph)

	finalState := (1 << size) - 1
	visited := make(map[[2]int]struct{})
	for i := 0; i < size; i++ {
		queue = append(queue, []int{i, 1 << i})
		visited[[2]int{i, 1 << i}] = struct{}{}
	}

	ans := -1
	for len(queue) > 0 {
		nq := make([][]int, 0)
		ans++
		for _, item := range queue {
			node := item[0]
			state := item[1]
			for _, next := range graph[node] {
				nextState := state | (1 << next)
				if nextState == finalState {
					return ans + 1
				}
				if _, ok := visited[[2]int{next, nextState}]; !ok {
					nq = append(nq, []int{next, nextState})
					visited[[2]int{next, nextState}] = struct{}{}
				}
			}
		}
		queue = nq
	}
	return 0
}
