package Solution

func Solution(n int) int {
	// dfs内存爆了，试试bfs
	in := make(map[int]struct{})
	queue := [][2]int{{n, 0}}
	for len(queue) > 0 {
		nq := make([][2]int, 0)
		for _, cur := range queue {
			if cur[0] == 1 {
				return cur[1] + 1
			}
			next := cur[0] - 1
			if _, ok := in[next]; !ok {
				in[next] = struct{}{}
				nq = append(nq, [2]int{next, cur[1] + 1})
			}
			if cur[0]%2 == 0 {
				next = cur[0] / 2
				if _, ok := in[next]; !ok {
					in[next] = struct{}{}
					nq = append(nq, [2]int{next, cur[1] + 1})
				}
			}

			if cur[0]%3 == 0 {
				next = cur[0] / 3
				if _, ok := in[next]; !ok {
					in[next] = struct{}{}
					nq = append(nq, [2]int{next, cur[1] + 1})
				}
			}
		}
		queue = nq
	}
	return -1
}
