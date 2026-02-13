package Solution

import "sort"

func Solution(n int, m int, hBars []int, vBars []int) int {
	// 能相邻，就会扩大最大相邻长度
	sort.Ints(hBars)
	sort.Ints(vBars)
	mh, mv := 1, 1
	cnt := 1

	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		mh = max(mh, cnt)
	}
	mh = max(mh, cnt)
	cnt = 1
	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		mv = max(mv, cnt)
	}
	mv = max(mv, cnt)
	edge := min(mv, mh) + 1
	return edge * edge
}
