package Solution

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func leastInterval(tasks []byte, n int) int {
	tmp := make([]int, 26)
	temp := []int{}
	p := 1
	for _, v := range tasks {
		tmp[v-'A']++
	}
	for _, v := range tmp {
		temp = append(temp, v)
	}
	sort.Slice(temp, func(a, b int) bool {
		return temp[a] > temp[b]
	})
	for i := 1; i < len(temp); i++ {
		if temp[0] == temp[i] {
			p++
		}
	}
	return max((n+1)*(temp[0]-1)+p, len(tasks))
}
