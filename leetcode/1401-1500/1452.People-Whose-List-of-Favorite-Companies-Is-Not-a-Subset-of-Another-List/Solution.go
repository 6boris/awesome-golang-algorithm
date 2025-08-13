package Solution

func Solution(favoriteCompanies [][]string) []int {
	l := len(favoriteCompanies)
	subMaps := make([]map[string]struct{}, l)
	for i := range favoriteCompanies {
		subMaps[i] = map[string]struct{}{}
		for _, com := range favoriteCompanies[i] {
			subMaps[i][com] = struct{}{}
		}
	}
	var ans []int
	for i := 0; i < l; i++ {
		ok := true
		for next := 0; next < l && ok; next++ {
			if next == i || len(subMaps[i]) > len(subMaps[next]) {
				continue
			}
			in := true
			for k := range subMaps[i] {
				if _, ok1 := subMaps[next][k]; !ok1 {
					in = false
					break
				}
			}
			ok = !in
		}
		if ok {
			ans = append(ans, i)
		}
	}
	return ans
}
