package Solution

import "sort"

type item2007 struct {
	indies []int
	idx    int
}

func Solution(changed []int) []int {
	n := len(changed)
	if n&1 == 1 {
		return []int{}
	}

	cnt := make(map[int]item2007)
	sort.Ints(changed)
	used := make([]bool, n)
	for i := range changed {
		v, ok := cnt[changed[i]]
		if !ok {
			cnt[changed[i]] = item2007{
				indies: []int{i},
				idx:    0,
			}
			continue
		}
		v.indies = append(v.indies, i)
		cnt[changed[i]] = v
	}

	var (
		target int
		ret    []int
	)

	// 1, 2, 3, 4, 6, 8
	// 0,0,0,0
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		sv := cnt[changed[i]]
		sv.idx++
		cnt[changed[i]] = sv
		target = changed[i] * 2
		v, ok := cnt[target]
		if !ok || v.idx >= len(v.indies) {
			return []int{}
		}
		ret = append(ret, changed[i])
		used[v.indies[v.idx]] = true
		v.idx++
		cnt[target] = v
	}

	return ret
}
