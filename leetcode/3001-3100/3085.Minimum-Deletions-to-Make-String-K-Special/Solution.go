package Solution

import "sort"

func Solution(word string, k int) int {
	count := [26]int{}
	for _, b := range word {
		count[b-'a']++
	}

	list := make([]int, 0)
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			list = append(list, count[i])
		}
	}
	ll := len(list)

	sum := make([]int, ll)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	sum[0] = list[0]
	for i := 1; i < ll; i++ {
		sum[i] = sum[i-1] + list[i]
	}

	ans := -1
	//26*26 N^2 ojbk
	for start := 0; start < ll; start++ {
		leftRemove := 0
		if start > 0 {
			leftRemove = sum[start-1]
		}

		for end := 0; end < ll; end++ {
			rightRemove := sum[ll-1] - sum[end]
			used := 0

			for try := end; try > start && list[try]-list[start] > k; try-- {
				used += list[try] - list[start] - k
			}
			if ans == -1 || rightRemove+leftRemove+used < ans {
				ans = rightRemove + leftRemove + used
			}
		}
	}

	return ans
}
