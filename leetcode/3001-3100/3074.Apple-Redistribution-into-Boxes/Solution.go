package Solution

import "sort"

func Solution(apple []int, capacity []int) int {
	sum := 0
	for _, n := range apple {
		sum += n
	}
	var ret int
	sort.Ints(capacity)
	for i := len(capacity) - 1; i >= 0 && sum > 0; i-- {
		sum -= capacity[i]
		ret++
	}
	return ret
}
