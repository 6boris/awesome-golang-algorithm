package Solution

func Solution(fronts []int, backs []int) int {
	fm := make(map[int]struct{})
	for _, n := range fronts {
		fm[n] = struct{}{}
	}
	ans := 0
	for _, n := range backs {
		if _, ok := fm[n]; !ok {
			if ans == 0 || ans > n {
				ans = n
			}
		}
	}
	mm := make(map[int][]int)
	keys := make([]int, 0)
	for i, f := range fronts {
		if _, ok := mm[f]; !ok {
			mm[f] = make([]int, 0)
			keys = append(keys, f)
		}
		mm[f] = append(mm[f], backs[i])
	}
	for i := 0; i < len(keys); i++ {
		ok := true
		for _, v := range mm[keys[i]] {
			if v == keys[i] {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		if ans == 0 || ans > keys[i] {
			ans = keys[i]
		}
	}
	return ans
}
