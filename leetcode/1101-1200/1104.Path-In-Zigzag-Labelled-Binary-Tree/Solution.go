package Solution

func Solution(label int) []int {
	res := make([]int, 0)
	idx, base := 0, 1
	for ; base <= label; base, idx = base*2, idx+1 {
	}

	base, idx = base/2, idx-1
	for {
		nextBase, nextIdx := base/2, idx-1
		res = append(res, label)
		if nextIdx < 0 {
			break
		}
		diff := (label - base) / 2
		label = base - 1 - diff
		base, idx = nextBase, nextIdx
	}

	for s, e := 0, len(res)-1; s < e; s, e = s+1, e-1 {
		res[s], res[e] = res[e], res[s]
	}

	return res
}
