package Solution

func Solution(responses [][]string) string {
	store := make(map[string]int)
	for _, resp := range responses {
		cur := make(map[string]struct{})
		for _, str := range resp {
			cur[str] = struct{}{}
		}
		for key := range cur {
			store[key]++
		}
	}
	var (
		ret string
		cnt int
	)
	for key, count := range store {
		if count > cnt {
			ret = key
			cnt = count
			continue
		}

		if count == cnt {
			ret = min(ret, key)
		}
	}

	return ret
}
