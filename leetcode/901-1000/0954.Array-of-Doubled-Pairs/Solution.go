package Solution

import "sort"

type index954 struct {
	idx    int
	indies []int
}

func Solution(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		a := arr[i]
		b := arr[j]
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		return a < b
	})

	cnt := make(map[int]index954)
	for i, n := range arr {
		if _, ok := cnt[n]; !ok {
			cnt[n] = index954{idx: 0, indies: []int{}}
		}
		ic := cnt[n]
		ic.indies = append(ic.indies, i)
		cnt[n] = ic
	}

	skip := make([]bool, len(arr))
	index := 0
	for ; index < len(arr) && arr[index] == 0; index++ {
	}
	if index&1 == 1 {
		return false
	}
	for ; index < len(arr); index++ {
		if skip[index] {
			continue
		}
		target := arr[index] * 2
		v, ok := cnt[target]
		if !ok {
			return false
		}
		if v.idx == len(v.indies) {
			return false
		}
		skip[v.indies[v.idx]] = true
		v.idx++
		cnt[target] = v
	}
	return true
}
