package Solution

import "sort"

func Solution(n int, logs [][]int, x int, queries []int) []int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	m := len(queries)
	idx := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool {
		return queries[idx[a]] < queries[idx[b]]
	})

	ln := len(logs)
	cnt := make([]int, n+1)
	distinct := 0
	l, r := 0, 0
	ret := make([]int, m)

	for _, qi := range idx {
		q := queries[qi]
		for r < ln && logs[r][1] <= q { // 加入 time <= q
			id := logs[r][0]
			if cnt[id] == 0 {
				distinct++
			}
			cnt[id]++
			r++
		}
		for l < r && logs[l][1] < q-x { // 移出 time < q-x
			id := logs[l][0]
			cnt[id]--
			if cnt[id] == 0 {
				distinct--
			}
			l++
		}
		ret[qi] = n - distinct
	}
	return ret
}
