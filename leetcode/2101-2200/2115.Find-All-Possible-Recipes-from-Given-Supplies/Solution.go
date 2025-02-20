package Solution

func Solution(recipes []string, ingredients [][]string, supplies []string) []string {
	sm := make(map[string]struct{})
	rm := make(map[string]struct{})
	for _, s := range supplies {
		sm[s] = struct{}{}
	}
	for _, r := range recipes {
		rm[r] = struct{}{}
	}

	in := make(map[string]int)
	unlock := make(map[string][]string)
	for idx, r := range recipes {
		in[r] = 0
		for _, i := range ingredients[idx] {
			if _, ok := sm[i]; ok {
				continue
			}
			if i == r {
				in[r] = -1
				break
			}
			if _, ok := rm[i]; ok {
				in[r]++
				if _, ok := unlock[i]; !ok {
					unlock[i] = []string{}
				}
				unlock[i] = append(unlock[i], r)
				continue
			}
			in[r] = -1 // 无法组成,用了一些不存在的东西
			break
		}
	}
	res := []string{}
	queue := []string{}
	for r, c := range in {
		if c == 0 {
			queue = append(queue, r)
			res = append(res, r)
		}
	}
	for len(queue) > 0 {
		nq := make([]string, 0)
		for _, cur := range queue {
			for _, r := range unlock[cur] {
				in[r]--
				if in[r] == 0 {
					nq = append(nq, r)
					res = append(res, r)
				}
			}
		}
		if len(nq) == 0 {
			break
		}
		queue = nq
	}

	return res
}
