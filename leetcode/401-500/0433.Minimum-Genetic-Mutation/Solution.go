package Solution

type Step433 struct {
	gene string
	step int
}

func Solution(startGene string, endGene string, bank []string) int {
	index := 0
	lb := len(bank)
	for ; index < lb; index++ {
		if bank[index] == endGene {
			break
		}
	}
	if index == lb {
		return -1
	}
	if startGene == endGene {
		return 0
	}

	mutFn := func(a, b string) bool {
		diff := 0
		for i := range 8 {
			if a[i] == b[i] {
				continue
			}
			diff++
			if diff > 1 {
				return false
			}
		}
		return diff == 1
	}
	mutations := make(map[string][]string)
	for i := 0; i < lb-1; i++ {
		for j := i + 1; j < lb; j++ {
			if mutFn(bank[i], bank[j]) {
				mutations[bank[i]] = append(mutations[bank[i]], bank[j])
				mutations[bank[j]] = append(mutations[bank[j]], bank[i])
			}
		}
	}
	for i := range lb {
		if mutFn(startGene, bank[i]) {
			mutations[startGene] = append(mutations[startGene], bank[i])
		}
	}

	queue := []Step433{
		{startGene, 0},
	}
	used := map[string]struct{}{
		startGene: struct{}{},
	}

	for len(queue) > 0 {
		nq := make([]Step433, 0)
		for i := range queue {
			rel := mutations[queue[i].gene]
			for j := range rel {
				if rel[j] == endGene {
					return queue[i].step + 1
				}
				if _, ok := used[rel[j]]; ok {
					continue
				}
				used[rel[j]] = struct{}{}
				nq = append(nq, Step433{rel[j], queue[i].step + 1})
			}
		}
		queue = nq
	}

	return -1
}
