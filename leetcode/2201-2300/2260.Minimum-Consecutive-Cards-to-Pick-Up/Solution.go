package Solution

func Solution(cards []int) int {
	pos := make(map[int]int)
	ret := -1
	var (
		index, diff int

		ok bool
	)
	for i := range cards {
		index, ok = pos[cards[i]]
		if !ok {
			pos[cards[i]] = i
			continue
		}

		diff = i - index + 1
		if ret == -1 || ret > diff {
			ret = diff
		}
		pos[cards[i]] = i
	}
	return ret
}
