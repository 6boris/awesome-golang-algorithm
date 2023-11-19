package Solution

import "sort"

func Solution(nums []int) int {
	count := make(map[int]int)
	distinct := make([]int, 0)
	for _, n := range nums {
		count[n]++
		if count[n] == 1 {
			distinct = append(distinct, n)
		}
	}

	sort.Ints(distinct)
	ans, sum := 0, 0
	for i := len(distinct) - 1; i > 0; i-- {
		sum += count[distinct[i]]
		ans += sum
	}

	return ans
}
