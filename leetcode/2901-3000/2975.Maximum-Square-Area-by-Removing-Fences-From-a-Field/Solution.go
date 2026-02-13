package Solution

import "sort"

func Solution(m int, n int, hFences []int, vFences []int) int {
	hEdges := getEdges(hFences, m)
	vEdges := getEdges(vFences, n)

	var res int64 = 0
	for e := range hEdges {
		if _, exists := vEdges[e]; exists {
			if int64(e) > res {
				res = int64(e)
			}
		}
	}

	if res == 0 {
		return -1
	}
	return int((res * res) % 1000000007)
}

func getEdges(fences []int, border int) map[int]bool {
	set := make(map[int]bool)
	list := make([]int, len(fences))

	copy(list, fences)
	list = append(list, 1, border)
	sort.Ints(list)

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			set[list[j]-list[i]] = true
		}
	}

	return set
}
