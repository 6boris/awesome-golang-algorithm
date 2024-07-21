package Solution

func sortConditions(k int, conditions [][]int) (map[int]int, bool) {
	children := make(map[int][]int)
	in := make(map[int]int)
	for _, c := range conditions {
		if _, ok := children[c[0]]; !ok {
			children[c[0]] = make([]int, 0)
		}
		children[c[0]] = append(children[c[0]], c[1])
		in[c[1]]++
	}
	index := 0
	pos := make(map[int]int)
	have := make([]int, 0)
	left := k
	for i := 1; i <= k; i++ {
		if _, ok := in[i]; !ok {
			pos[i] = index
			have = append(have, i)
			index++
			left--
		}
	}
	for left > 0 {
		next := make([]int, 0)
		for _, cur := range have {
			for _, rel := range children[cur] {
				in[rel]--
				if in[rel] == 0 {
					next = append(next, rel)
					left--
					pos[rel] = index
					index++
				}
			}
		}
		if len(next) == 0 && left > 0 {
			return nil, false
		}
		have = next
	}
	return pos, true
}

func Solution(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = make([]int, k)
	}
	rowPos, ok := sortConditions(k, rowConditions)
	if !ok {
		return nil
	}

	colPos, ok := sortConditions(k, colConditions)
	if !ok {
		return nil
	}

	for i := 1; i <= k; i++ {
		ans[rowPos[i]][colPos[i]] = i
	}

	return ans
}
